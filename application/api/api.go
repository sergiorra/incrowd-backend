package api

import (
	"context"
	"fmt"
	"incrowd-backend/application/api/handlers"
	"incrowd-backend/application/api/handlers/article_detail"
	"incrowd-backend/application/api/handlers/articles_list"
	"incrowd-backend/domain/services/article_detail_service"
	"incrowd-backend/domain/services/articles_list_service"
	"incrowd-backend/internal/context_wrapper"
	"net"
	"strconv"
	"time"

	"incrowd-backend/application/api/middlewares"
	"incrowd-backend/config"
	"incrowd-backend/domain/ports/database"
	"incrowd-backend/internal/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	server *echo.Echo
	config config.Config
	Addr   string
}

func NewApi(articleRepository database.ArticleRepository, config config.Config) *API {
	return &API{
		server: echoServer(articleRepository, config),
		config: config,
		Addr:   apiAddr(config.Api),
	}
}

func (a *API) Start() error {
	log.Infof("HTTP listener running on %s", a.Addr)
	return a.server.Start(a.Addr)
}

// Shutdown gracefully shuts down the API server
func (a *API) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(a.config.App.ShutdownTimeoutInSeconds)*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown API server gracefully with error %s", err)
	}
}

// echoServer sets up an Echo server with various middlewares for handling HTTP requests
func echoServer(articleRepository database.ArticleRepository, config config.Config) *echo.Echo {
	e := echo.New()

	e.Logger.SetLevel(log.Lvl(config.App.LogLevel))

	e.Server.ReadHeaderTimeout = time.Duration(config.Api.ReadHeaderTimeoutInSeconds) * time.Second

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper: middleware.DefaultSkipper,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Errort(context_wrapper.GetCorrelationID(c.Request().Context()), fmt.Sprintf("[PANIC RECOVER] Error: %s", err))
			log.Errort(context_wrapper.GetCorrelationID(c.Request().Context()), fmt.Sprintf("[PANIC RECOVER] Stack trace: %s", string(stack)))
			return err
		},
	}))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(config.Api.TimeoutInSeconds) * time.Second,
	}))

	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Output: e.Logger.Output(),
			Format: `{"time":"${time_rfc3339_nano}","trackID":"${id}","remote_ip":"${remote_ip}",` +
				`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
				`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
				`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		}))

	e.Use(middlewares.CorrelationID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.GET("healthcheck", func(c echo.Context) error {
		return handlers.Healthcheck(c)
	})

	adh := article_detail.NewArticleDetailHandler(article_detail_service.NewArticleDetailService(articleRepository))
	e.GET("/provider/realise/v1/teams/:teamID/news/:articleID", func(c echo.Context) error {
		return adh.ArticleDetail(c)
	})

	alh := articles_list.NewArticlesListHandler(articles_list_service.NewArticlesListService(articleRepository))
	e.GET("/provider/realise/v1/teams/:teamID/news", func(c echo.Context) error {
		return alh.ArticlesList(c)
	})

	return e
}

func apiAddr(cfg config.Api) string {
	return net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
}

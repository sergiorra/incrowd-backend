package main

import (
	"context"
	"flag"
	"incrowd-backend/application/api"
	"incrowd-backend/infrastructure/providers/huddersfield_town"
	"incrowd-backend/internal/httpclient"
	"net/http"
	"os/signal"
	"syscall"

	"incrowd-backend/application/worker/article_poller_handler"
	"incrowd-backend/domain/services/article_poller_service"
	"incrowd-backend/infrastructure/database"
	"incrowd-backend/infrastructure/database/article"

	"incrowd-backend/config"
	"incrowd-backend/internal/log"
)

func main() {
	configFile := flag.String("conf", "config/config.local.json", "Config file path")
	flag.Parse()

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	appConfig, err := config.Read(*configFile)
	if err != nil {
		log.Fatalf("could not read config file with error %s", err)
	}

	log.SetupLogging(appConfig.App.LogLevel)

	dbClient, err := database.SetupDatabaseConnection(appConfig.Database)
	if err != nil {
		log.Fatalf("could not setup database connection with error %s", err)
	}

	defer func() {
		if err = dbClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("could not disconnect DB client with error %s", err)
		}
	}()

	httpClient := httpclient.NewHttpClient(appConfig.Http)

	articleRepository := article.NewArticleRespository(dbClient)
	huddersfieldTownProvider := huddersfield_town.NewHuddersfieldTownProvider(appConfig.HuddersfieldTownProvider, httpClient)

	articlePollerService := article_poller_service.NewArticlePollerService(articleRepository, huddersfieldTownProvider)
	articlePollerHandler := article_poller_handler.NewArticlePollerHandler(articlePollerService, appConfig.ArticlePoller)

	incrowdAPI := api.NewApi(articleRepository, *appConfig)

	go runArticlePollerScheduler(articlePollerHandler)
	go runApiHandler(incrowdAPI)

	<-shutdownCtx.Done()
	articlePollerHandler.Shutdown()
	incrowdAPI.Shutdown()
}

func runApiHandler(incrowdAPI *api.API) {
	if err := incrowdAPI.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start API server with error %s", err)
	}
}

func runArticlePollerScheduler(articlePollerHandler *article_poller_handler.ArticlePollerHandler) {
	if err := articlePollerHandler.HandleArticlePolling(); err != nil {
		log.Fatalf("could not schedule article poller with error %s", err)
	}
}

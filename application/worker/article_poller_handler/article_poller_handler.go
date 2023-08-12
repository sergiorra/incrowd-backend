package article_poller_handler

import (
	"context"
	"fmt"
	"time"

	"incrowd-backend/config"
	"incrowd-backend/domain/services/article_poller_service"
	"incrowd-backend/internal/context_wrapper"
	"incrowd-backend/internal/log"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
)

type ArticlePollerHandler struct {
	articlePollerService *article_poller_service.ArticlePollerService
	config               config.ArticlePoller
	scheduler            *gocron.Scheduler
}

func NewArticlePollerHandler(articlePollerService *article_poller_service.ArticlePollerService, config config.ArticlePoller) *ArticlePollerHandler {
	return &ArticlePollerHandler{
		articlePollerService: articlePollerService,
		config:               config,
		scheduler:            gocron.NewScheduler(time.UTC),
	}
}

// HandleArticlePolling handles periodic article polling using a scheduler
func (h *ArticlePollerHandler) HandleArticlePolling() error {
	_, err := h.scheduler.Every(h.config.ExecutionIntervalInMinutes).Minutes().Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.ExecutionTimeoutInSeconds)*time.Second)
		defer cancel()

		correlationID := uuid.New().String()
		ctx = context_wrapper.WithCorrelationID(ctx, correlationID)

		log.Infot(context_wrapper.GetCorrelationID(ctx), "Starting to run article poller")

		if err := h.articlePollerService.HandleArticlePolling(ctx); err != nil {
			log.Errort(context_wrapper.GetCorrelationID(ctx), fmt.Sprintf("could not run article polling with error %s", err))
		}
	})

	h.scheduler.StartAsync()

	if err != nil {
		return err
	}

	return nil
}

func (h *ArticlePollerHandler) Shutdown() {
}

package article_poller_service

import (
	"context"
	"fmt"
	"incrowd-backend/domain/models"
	"incrowd-backend/domain/ports/providers"
	"incrowd-backend/internal/context_wrapper"
	"incrowd-backend/internal/log"

	"incrowd-backend/domain/ports/database"
)

type ArticlePollerService struct {
	articleRepository        database.ArticleRepository
	huddersfieldTownProvider providers.HuddersfieldTownProvider
}

func NewArticlePollerService(articleRepository database.ArticleRepository, huddersfieldTownProvider providers.HuddersfieldTownProvider) *ArticlePollerService {
	return &ArticlePollerService{
		articleRepository:        articleRepository,
		huddersfieldTownProvider: huddersfieldTownProvider,
	}
}

// HandleArticlePolling polls for new articles, retrieves their details, and upserts them into the database
func (s *ArticlePollerService) HandleArticlePolling(ctx context.Context) error {
	articlesIDs, err := s.huddersfieldTownProvider.GetNewArticlesIDs(ctx)
	if err != nil {
		return fmt.Errorf("could not get new articles IDs with error %s", err)
	}

	articles := make([]*models.Article, 0, len(articlesIDs))

	for _, articleID := range articlesIDs {
		articleDetails, err := s.huddersfieldTownProvider.GetArticleInformation(ctx, articleID)
		if err != nil {
			log.Errort(context_wrapper.GetCorrelationID(ctx), fmt.Sprintf("could not get article information for id %s, with error %s", articleID, err))
			continue
		}

		articles = append(articles, articleDetails)
	}

	if len(articles) == 0 {
		return nil
	}

	if err = s.articleRepository.Upsert(ctx, articles[0].TeamID, articles); err != nil {
		return fmt.Errorf("could not execute upsert operation with error %s", err)
	}

	return nil
}

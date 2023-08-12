package providers

import (
	"context"

	"incrowd-backend/domain/models"
)

type HuddersfieldTownProvider interface {
	GetNewArticlesIDs(ctx context.Context) ([]string, error)
	GetArticleInformation(ctx context.Context, id string) (*models.Article, error)
}

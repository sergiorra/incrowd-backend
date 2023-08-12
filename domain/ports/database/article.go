package database

import (
	"context"

	"incrowd-backend/domain/models"
)

type ArticleRepository interface {
	GetByTeamIDAndID(ctx context.Context, teamID, ID string) (*models.Article, error)
	ListByTeamID(ctx context.Context, md models.MetaData, teamID string) ([]*models.Article, error)
	Upsert(ctx context.Context, teamID string, articles []*models.Article) error
	GetCollectionNames(ctx context.Context) ([]string, error)
}

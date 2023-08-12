package articles_list_service

import (
	"context"
	"incrowd-backend/domain/models"
	"incrowd-backend/internal/common"

	"incrowd-backend/domain/ports/database"
)

type ArticlesListService struct {
	articleRepository database.ArticleRepository
}

func NewArticlesListService(articleRepository database.ArticleRepository) *ArticlesListService {
	return &ArticlesListService{
		articleRepository: articleRepository,
	}
}

// ArticlesList retrieves a list of articles for a specified team ID, applying pagination and sorting sent in the metadata
func (s *ArticlesListService) ArticlesList(ctx context.Context, teamID string, md models.MetaData) ([]*models.Article, error) {
	collectionNames, err := s.articleRepository.GetCollectionNames(ctx)
	if err != nil {
		return nil, models.ErrInternalServer{}
	}

	if !common.Contains(collectionNames, teamID) {
		return nil, models.ErrTeamIDNotFound
	}

	articles, err := s.articleRepository.ListByTeamID(ctx, md, teamID)
	if err != nil {
		return nil, models.ErrInternalServer{}
	}

	return articles, nil
}

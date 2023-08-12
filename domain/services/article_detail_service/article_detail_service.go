package article_detail_service

import (
	"context"
	"incrowd-backend/domain/models"
	"incrowd-backend/domain/ports/database"
	"incrowd-backend/internal/common"
)

type ArticleDetailService struct {
	articleRepository database.ArticleRepository
}

func NewArticleDetailService(articleRepository database.ArticleRepository) *ArticleDetailService {
	return &ArticleDetailService{
		articleRepository: articleRepository,
	}
}

// ArticleDetail retrieves an article's details based on the provided team ID and article ID
func (s *ArticleDetailService) ArticleDetail(ctx context.Context, teamID, articleID string) (*models.Article, error) {
	collectionNames, err := s.articleRepository.GetCollectionNames(ctx)
	if err != nil {
		return nil, models.ErrInternalServer{}
	}

	if !common.Contains(collectionNames, teamID) {
		return nil, models.ErrTeamIDNotFound
	}

	article, err := s.articleRepository.GetByTeamIDAndID(ctx, teamID, articleID)
	if err != nil {
		switch err {
		case models.ErrArticleIDNotFound:
			return nil, err
		default:
			return nil, models.ErrInternalServer{}
		}
	}

	return article, nil
}

package article

import (
	"incrowd-backend/domain/models"
)

func mapArticleToMongoDbModel(a *models.Article) *Article {
	return &Article{
		ID:          a.ID,
		TeamID:      a.TeamID,
		OptaMatchID: a.OptaMatchID,
		Title:       a.Title,
		Teaser:      a.Teaser,
		Content:     a.Content,
		URL:         a.URL,
		VideoURL:    a.VideoURL,
		Published:   a.Published,
	}
}

func mapArticleListToDomainModel(articles []Article) []*models.Article {
	result := make([]*models.Article, 0, len(articles))

	for _, a := range articles {
		result = append(result, &models.Article{
			ID:          a.ID,
			TeamID:      a.TeamID,
			OptaMatchID: a.OptaMatchID,
			Title:       a.Title,
			Teaser:      a.Teaser,
			Content:     a.Content,
			URL:         a.URL,
			VideoURL:    a.VideoURL,
			Published:   a.Published,
		})
	}

	return result
}

func mapArticleToDomainModel(a *Article) *models.Article {
	return &models.Article{
		ID:          a.ID,
		TeamID:      a.TeamID,
		OptaMatchID: a.OptaMatchID,
		Title:       a.Title,
		Teaser:      a.Teaser,
		Content:     a.Content,
		URL:         a.URL,
		VideoURL:    a.VideoURL,
		Published:   a.Published,
	}
}

package article_detail

import (
	"time"

	"incrowd-backend/domain/models"
)

const (
	responseDateLayout = "2006-01-02T15:04:05.999Z"
)

func convertResponse(article *models.Article) *ArticleDetailResponse {
	return &ArticleDetailResponse{
		Status: "success",
		Data: ArticleData{
			ID:          article.ID,
			TeamID:      article.TeamID,
			OptaMatchID: article.OptaMatchID,
			Title:       article.Title,
			Teaser:      article.Teaser,
			Content:     article.Content,
			URL:         article.URL,
			VideoURL:    article.VideoURL,
			Published:   article.Published.Format(responseDateLayout),
		},
		Metadata: Metadata{
			CreatedAt: time.Now().Format(responseDateLayout),
		},
	}
}

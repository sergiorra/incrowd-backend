package articles_list

import (
	"time"

	"incrowd-backend/domain/models"
)

const (
	responseDateLayout = "2006-01-02T15:04:05.999Z"
)

func convertResponse(md models.MetaData, articles []*models.Article) *ArticlesListResponse {
	return &ArticlesListResponse{
		Status: "success",
		Data:   mapArticlesData(articles),
		Metadata: Metadata{
			CreatedAt: time.Now().Format(responseDateLayout),
			Count:     md.Count,
			Page:      md.Page,
			Sort:      md.Sort,
			Order:     md.Order,
		},
	}
}

func mapArticlesData(articles []*models.Article) []Article {
	result := make([]Article, 0, len(articles))

	for _, article := range articles {
		a := Article{
			ID:          article.ID,
			TeamID:      article.TeamID,
			OptaMatchID: article.OptaMatchID,
			Title:       article.Title,
			Teaser:      article.Teaser,
			Content:     article.Content,
			URL:         article.URL,
			VideoURL:    article.VideoURL,
			Published:   article.Published.Format(responseDateLayout),
		}

		result = append(result, a)
	}

	return result
}

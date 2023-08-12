package huddersfield_town

import (
	"context"

	"incrowd-backend/config"
	"incrowd-backend/domain/models"
	"incrowd-backend/internal/common"
)

const (
	publishDateLayout = "2006-01-02 15:04:05"
)

func (r GetNewArticlesIDsResponse) mapToIDs() []string {
	result := make([]string, 0, len(r.NewsletterNewsItems))

	for _, v := range r.NewsletterNewsItems {
		result = append(result, v.NewsArticleID)
	}

	return result
}

func (r GetArticleInformationResponse) mapToArticleModel(ctx context.Context, config config.Provider) *models.Article {
	return &models.Article{
		ID:          r.NewsArticle.NewsArticleID,
		TeamID:      config.TeamID,
		OptaMatchID: r.NewsArticle.OptaMatchId,
		Title:       r.NewsArticle.Title,
		Teaser:      r.NewsArticle.TeaserText,
		Content:     r.NewsArticle.BodyText,
		URL:         r.NewsArticle.ArticleURL,
		VideoURL:    r.NewsArticle.VideoURL,
		Published:   common.ConvertStringToDate(ctx, publishDateLayout, r.NewsArticle.PublishDate),
	}
}

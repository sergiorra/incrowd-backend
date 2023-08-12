package huddersfield_town

type GetNewArticlesIDsResponse struct {
	NewsletterNewsItems []NewArticlesListItem `xml:"NewsletterNewsItems>NewsletterNewsItem"`
}

type NewArticlesListItem struct {
	NewsArticleID string `xml:"NewsArticleID"`
}

type GetArticleInformationResponse struct {
	NewsArticle ArticleInformation `xml:"NewsArticle"`
}

type ArticleInformation struct {
	NewsArticleID string `xml:"NewsArticleID"`
	OptaMatchId   string `xml:"OptaMatchId"`
	Title         string `xml:"Title"`
	TeaserText    string `xml:"TeaserText"`
	BodyText      string `xml:"BodyText"`
	ArticleURL    string `xml:"ArticleURL"`
	VideoURL      string `xml:"VideoURL"`
	PublishDate   string `xml:"PublishDate"`
}

package article_detail

type ArticleDetailResponse struct {
	Status   string      `json:"status"`
	Data     ArticleData `json:"data"`
	Metadata Metadata    `json:"metadata"`
}

type ArticleData struct {
	ID          string `json:"id"`
	TeamID      string `json:"teamId"`
	OptaMatchID string `json:"optaMatchId"`
	Title       string `json:"title"`
	Teaser      string `json:"teaser"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	VideoURL    string `json:"videoUrl"`
	Published   string `json:"published"`
}

type Metadata struct {
	CreatedAt string `json:"createdAt"`
}

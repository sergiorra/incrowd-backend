package articles_list

type ArticlesListResponse struct {
	Status   string    `json:"status"`
	Data     []Article `json:"data"`
	Metadata Metadata  `json:"metadata"`
}

type Article struct {
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
	Count     int    `json:"count"`
	Page      int    `json:"page"`
	Sort      string `json:"sort"`
	Order     string `json:"order"`
}

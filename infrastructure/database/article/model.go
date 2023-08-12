package article

import "time"

type Article struct {
	ID          string    `bson:"_id"`
	TeamID      string    `bson:"teamId"`
	OptaMatchID string    `bson:"optaMatchId,omitempty"`
	Title       string    `bson:"title"`
	Teaser      string    `bson:"teaser,omitempty"`
	Content     string    `bson:"content,omitempty"`
	URL         string    `bson:"url,omitempty"`
	VideoURL    string    `bson:"videoUrl,omitempty"`
	Published   time.Time `bson:"published"`
}

package models

import "time"

type Article struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"teamId"`
	OptaMatchID string    `json:"optaMatchId"`
	Title       string    `json:"title"`
	Teaser      string    `json:"teaser"`
	Content     string    `json:"content"`
	URL         string    `json:"url"`
	VideoURL    string    `json:"videoUrl"`
	Published   time.Time `json:"published"`
}

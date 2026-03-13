package models

type AnimeDetail struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"coverImage"`
	Episodes    int    `json:"episodes"`
}

package models

type AniListDetailResponse struct {
	Data struct {
		Media struct {
			ID    int `json:"id"`
			Title struct {
				Romaji string `json:"romaji"`
			} `json:"title"`
			CoverImage struct {
				Large string `json:"large"`
			} `json:"coverImage"`
			Description string `json:"description"`
			Episodes    int    `json:"episodes"`
		} `json:"media"`
	} `json:"data"`
}

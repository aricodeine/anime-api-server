package models

type JikanDetailResponse struct {
	Data struct {
		ID     int    `json:"mal_id"`
		Title  string `json:"title_english"`
		Images struct {
			JPG struct {
				Large string `json:"large_image_url"`
			} `json:"jpg"`
		} `json:"images"`
		Description string `json:"synopsis"`
		Episodes    int    `json:"episodes"`
	} `json:"data"`
}

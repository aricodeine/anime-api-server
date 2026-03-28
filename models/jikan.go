package models

type JikanResponse struct {
	Data []struct {
		ID     int    `json:"mal_id"`
		Title  string `json:"title_english"`
		Images struct {
			JPG struct {
				Large string `json:"large_image_url"`
			} `json:"jpg"`
		} `json:"images"`
	} `json:"data"`
}

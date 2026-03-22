package jikan

import (
	"anime-server/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const jikanURL = "https://api.jikan.moe/v4/anime"
const jikanTimeout = 5 * time.Second

func SearchAnime(ctx context.Context, query string, page int, limit int) ([]models.Anime, error) {

	params := url.Values{}
	params.Set("q", query)
	params.Set("page", fmt.Sprintf("%d", page))
	params.Set("limit", fmt.Sprintf("%d", limit))

	ctx, cancel := context.WithTimeout(ctx, jikanTimeout)
	defer cancel()

	fullURL := fmt.Sprintf("%s?%s", jikanURL, params.Encode())
	req, _ := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jikanResp := models.JikanResponse{}
	animes := []models.Anime{}

	data, err := io.ReadAll(resp.Body)
	if err != nil || json.Unmarshal(data, &jikanResp) != nil {
		return nil, err
	}

	for _, m := range jikanResp.Data {
		animes = append(animes, models.Anime{
			ID:    m.ID,
			Title: m.Title,
			Cover: m.Images.JPG.Large,
		})
	}
	return animes, nil
}

func GetAnimeByID(ctx context.Context, id int) (*models.AnimeDetail, error) {

	params := url.Values{}
	params.Set("id", fmt.Sprintf("%d", id))

	ctx, cancel := context.WithTimeout(ctx, jikanTimeout)
	defer cancel()

	fullURL := fmt.Sprintf("%s/%d", jikanURL, id)

	req, _ := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 🎯 Mapping step
	aniResp := models.JikanDetailResponse{}
	data, err := io.ReadAll(resp.Body)
	if err != nil || json.Unmarshal(data, &aniResp) != nil {
		return nil, err
	}

	animeDetail := &models.AnimeDetail{
		ID:          aniResp.Data.ID,
		Title:       aniResp.Data.Title,
		Cover:       aniResp.Data.Images.JPG.Large,
		Description: aniResp.Data.Description,
		Episodes:    aniResp.Data.Episodes,
	}

	log.Printf("%s", animeDetail.Description)

	return animeDetail, nil
}

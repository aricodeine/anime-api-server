package anilist

import (
	"anime-server/models"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const anilistURL = "https://graphql.anilist.co"
const anilistTimeout = 5 * time.Second

func SearchAnime(ctx context.Context, query string, page int, limit int) ([]models.Anime, error) {
	gql := `
	query ($search: String, $page: Int, $perPage: Int) {
	  Page(page: $page, perPage: $perPage) {
		media(search: $search, type: ANIME) {
		  id
		  title {
			romaji
		  }
		  coverImage {
			large
		  }
		}
	  }
	}`

	body := map[string]any{
		"query": gql,
		"variables": map[string]any{
			"search":  query,
			"page":    page,
			"perPage": limit,
		},
	}

	jsonBody, _ := json.Marshal(body)

	ctx, cancel := context.WithTimeout(ctx, anilistTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", anilistURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 🎯 Mapping step
	aniResp := models.AniListResponse{}
	animes := []models.Anime{}

	data, err := io.ReadAll(resp.Body)
	if err != nil || json.Unmarshal(data, &aniResp) != nil {
		return nil, err
	}

	for _, m := range aniResp.Data.Page.Media {
		animes = append(animes, models.Anime{
			ID:    m.ID,
			Title: m.Title.Romaji,
			Cover: m.CoverImage.Large,
		})
	}
	return animes, nil
}

func GetAnimeByID(ctx context.Context, id int) (*models.AnimeDetail, error) {
	gql := `
	query($id: Int) {
  		Media(id: $id, type: ANIME) {
			id
			coverImage {
				large
			}
			title {
				romaji
			}
			description
			episodes
		}
	}`

	body := map[string]any{
		"query": gql,
		"variables": map[string]int{
			"id": id,
		},
	}

	jsonBody, _ := json.Marshal(body)

	ctx, cancel := context.WithTimeout(ctx, anilistTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", anilistURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 🎯 Mapping step
	aniResp := models.AniListDetailResponse{}
	data, err := io.ReadAll(resp.Body)
	if err != nil || json.Unmarshal(data, &aniResp) != nil {
		return nil, err
	}

	animeDetail := &models.AnimeDetail{
		ID:          aniResp.Data.Media.ID,
		Title:       aniResp.Data.Media.Title.Romaji,
		Cover:       aniResp.Data.Media.CoverImage.Large,
		Description: aniResp.Data.Media.Description,
		Episodes:    aniResp.Data.Media.Episodes,
	}

	return animeDetail, nil
}

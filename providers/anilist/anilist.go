package anilist

import (
	"anime-server/internal/cache"
	"anime-server/models"
	"anime-server/services"
	"context"
	"fmt"
	"time"
)

type AniListProvider struct {
	cache *cache.Cache
}

func New(cache *cache.Cache) *AniListProvider {
	return &AniListProvider{
		cache: cache,
	}
}

func (a *AniListProvider) Search(ctx context.Context, query string, page int, limit int) ([]models.Anime, error) {
	key := fmt.Sprintf("search:%s:%d:%d", query, page, limit)

	if cached, ok := a.cache.Get(key); ok {
		return cached.([]models.Anime), nil
	}

	result, err := services.SearchAnime(ctx, query, page, limit)
	if err != nil {
		return nil, err
	}
	a.cache.Set(key, result, 5*time.Minute)
	return result, nil
}

func (a *AniListProvider) GetByID(ctx context.Context, id int) (*models.AnimeDetail, error) {
	key := fmt.Sprintf("anime:%d", id)

	if cached, ok := a.cache.Get(key); ok {
		return cached.(*models.AnimeDetail), nil
	}

	result, err := services.GetAnimeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	a.cache.Set(key, result, 5*time.Minute)
	return result, nil
}

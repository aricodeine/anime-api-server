package anilist

import (
	"anime-server/internal/cache"
	"anime-server/models"
	"anime-server/services/anilist"
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

type AniListProvider struct {
	cache *cache.Cache
	group singleflight.Group
}

func New(cache *cache.Cache) *AniListProvider {
	return &AniListProvider{
		cache: cache,
	}
}

func (a *AniListProvider) Search(ctx context.Context, query string, page int, limit int) ([]models.Anime, error) {
	key := fmt.Sprintf("search:anilist:%s:%d:%d", query, page, limit)

	if cached, ok := a.cache.Get(key); ok {
		return cached.([]models.Anime), nil
	}

	result, err, _ := a.group.Do(key, func() (interface{}, error) {
		data, err := anilist.SearchAnime(ctx, query, page, limit)
		if err != nil {
			return nil, err
		}
		a.cache.Set(key, data, 5*time.Minute)
		return data, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]models.Anime), nil

}

func (a *AniListProvider) GetByID(ctx context.Context, id int) (*models.AnimeDetail, error) {
	key := fmt.Sprintf("anime:anilist:%d", id)

	if cached, ok := a.cache.Get(key); ok {
		return cached.(*models.AnimeDetail), nil
	}

	result, err, _ := a.group.Do(key, func() (interface{}, error) {
		data, err := anilist.GetAnimeByID(ctx, id)
		if err != nil {
			return nil, err
		}
		a.cache.Set(key, data, 5*time.Minute)
		return data, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*models.AnimeDetail), nil
}

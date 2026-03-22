package jikan

import (
	"anime-server/internal/cache"
	"anime-server/models"
	"anime-server/services/jikan"
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

type JikanProvider struct {
	cache *cache.Cache
	group singleflight.Group
}

func New(cache *cache.Cache) *JikanProvider {
	return &JikanProvider{
		cache: cache,
	}
}

func (a *JikanProvider) Search(ctx context.Context, query string, page int, limit int) ([]models.Anime, error) {
	key := fmt.Sprintf("search:%s:%d:%d", query, page, limit)

	if cached, ok := a.cache.Get(key); ok {
		return cached.([]models.Anime), nil
	}

	result, err, _ := a.group.Do(key, func() (interface{}, error) {
		data, err := jikan.SearchAnime(ctx, query, page, limit)
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

func (a *JikanProvider) GetByID(ctx context.Context, id int) (*models.AnimeDetail, error) {
	key := fmt.Sprintf("anime:%d", id)

	if cached, ok := a.cache.Get(key); ok {
		return cached.(*models.AnimeDetail), nil
	}

	result, err, _ := a.group.Do(key, func() (interface{}, error) {
		data, err := jikan.GetAnimeByID(ctx, id)
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

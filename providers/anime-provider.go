package providers

import (
	"anime-server/models"
	"context"
)

type AnimeProvider interface {
	Search(ctx context.Context, query string, page int, limit int) ([]models.Anime, error)
	GetByID(ctx context.Context, id int) (*models.AnimeDetail, error)
}

// Any provider that gives anime data must follow this

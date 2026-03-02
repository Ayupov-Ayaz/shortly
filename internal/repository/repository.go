package repository

import (
	"context"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
)

type Repository interface {
	// save originalURL and shortURL during expire
	Save(
		ctx context.Context,
		req *gen.CreateURLResponse,
	) error
	// get by origin url
	GetByOrigin(
		ctx context.Context,
		shortURL string,
	) (*gen.CreateURLResponse, error)
	// get by short url
	GetByShortURL(
		ctx context.Context,
		originalURL string,
	) (*gen.CreateURLResponse, error)
	// delete url
	Del(
		ctx context.Context,
		shortURL string,
	) error
}

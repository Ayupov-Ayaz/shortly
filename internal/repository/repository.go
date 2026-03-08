package repository

import (
	"context"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
)

type Repository interface {
	// create originalURL and shortURL during expire
	Create(
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
	Delete(
		ctx context.Context,
		shortURL string,
	) error
	//
	Update(
		ctx context.Context,
		req *gen.CreateURLResponse,
	) error
}

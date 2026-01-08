package storage

import (
	"context"

	"github.com/ayupov-ayaz/shortly/internal/entity"
)

type Storage interface {
	// save originalURL and shortURL during expire
	Save(
		ctx context.Context,
		req *entity.URLResponse,
	) error
	// get by origin url
	GetByOrigin(
		ctx context.Context,
		shortURL string,
	) (*entity.URLResponse, error)
	// get by short url
	GetByShortURL(
		ctx context.Context,
		originalURL string,
	) (*entity.URLResponse, error)
	// delete url
	Del(
		ctx context.Context,
		shortURL string,
	) error
}

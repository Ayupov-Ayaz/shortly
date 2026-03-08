package repository

import (
	"context"

	"github.com/ayupov-ayaz/shortly/internal/entity"
)

type Repository interface {
	// create originalURL and shortURL during expire
	Create(
		ctx context.Context,
		req entity.URL,
	) error
	// get by origin url
	GetByOrigin(
		ctx context.Context,
		shortURL string,
	) (*entity.URL, error)
	// get by short url
	GetByShortURL(
		ctx context.Context,
		originalURL string,
	) (*entity.URL, error)
	// delete url
	Delete(
		ctx context.Context,
		shortURL string,
	) error
	//
	Update(
		ctx context.Context,
		req entity.URL,
	) error
}

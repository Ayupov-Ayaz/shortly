package shortener

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/repository"
)

type Shortener interface {
	ShortenURL(
		ctx context.Context,
		req gen.CreateURLRequest,
	) (*gen.CreateURLResponse, error)
}

type idGenerator interface {
	Generate() string
}

var _ Shortener = (*URLShortener)(nil)

type URLShortener struct {
	storage   repository.Repository
	generator idGenerator

	baseURL       *url.URL //todo: use baseURL
	defaultExpire time.Duration
}

func New(
	storage repository.Repository,
	generator idGenerator,
	baseURL *url.URL,
	defaultExpire time.Duration,
) *URLShortener {
	return &URLShortener{
		storage:       storage,
		generator:     generator,
		baseURL:       baseURL,
		defaultExpire: defaultExpire,
	}
}

func (u *URLShortener) ShortenURL(
	ctx context.Context,
	req gen.CreateURLRequest,
) (*gen.CreateURLResponse, error) {
	// if original url already exists, return it
	resp, err := u.storage.GetByOrigin(ctx, req.Url)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	} else if resp != nil {
		return resp, nil
	}

	shortCode := u.generator.Generate()

	res := &gen.CreateURLResponse{
		CreatedAt:   time.Now(),
		ExpiresAt:   req.ExpiresAt,
		OriginalURL: req.Url,
		ShortURL:    shortCode,
	}

	// todo: use entity
	err = u.storage.Create(ctx, res)
	if err != nil {
		return nil, err
	}

	// todo: join base url and short code
	return res, nil
}

package shortener

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/entity"
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

	baseURL *url.URL
	now     func() time.Time
}

func New(
	storage repository.Repository,
	generator idGenerator,
	baseURL *url.URL,
	now func() time.Time,
) *URLShortener {
	return &URLShortener{
		storage:   storage,
		generator: generator,
		baseURL:   baseURL,
		now:       now,
	}
}

func (u *URLShortener) ShortenURL(
	ctx context.Context,
	req gen.CreateURLRequest,
) (*gen.CreateURLResponse, error) {
	// if original url already exists, return it
	entityURL, err := u.storage.GetByOrigin(ctx, req.Url)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	} else if entityURL != nil {
		return castEntityToAPIResp(*entityURL, u.baseURL), nil
	}

	shortCode := u.generator.Generate()

	shortURL := entity.URL{
		OriginalURL: req.Url,
		ShortCode:   shortCode,
		CreatedAt:   u.now(),
		ExpiresAt:   req.ExpiresAt,
	}

	err = u.storage.Create(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	res := castEntityToAPIResp(shortURL, u.baseURL)

	return res, nil
}

func castEntityToAPIResp(
	url entity.URL,
	baseURL *url.URL,
) *gen.CreateURLResponse {
	return &gen.CreateURLResponse{
		OriginalURL: url.OriginalURL,
		ShortURL:    baseURL.JoinPath(url.ShortCode).String(),
		CreatedAt:   url.CreatedAt,
		ExpiresAt:   url.ExpiresAt,
	}
}

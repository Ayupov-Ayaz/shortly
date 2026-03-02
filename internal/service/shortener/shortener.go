package shortener

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/repository"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
)

type Shortener interface {
	ShortenURL(
		ctx context.Context,
		req gen.CreateURLRequest,
	) (*gen.CreateURLResponse, error)
}

var _ Shortener = (*URLShortener)(nil)

type URLShortener struct {
	storage   repository.Repository
	generator id.Generator

	baseURL       *url.URL //todo: use baseURL
	defaultExpire time.Duration
}

func New(
	storage repository.Repository,
	generator id.Generator,
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

	shortURL := u.baseURL.JoinPath(u.generator.Generate())

	res := &gen.CreateURLResponse{
		CreatedAt:   time.Now(),
		ExpiresAt:   req.ExpiresAt,
		OriginalURL: req.Url,
		ShortURL:    shortURL.String(),
	}

	err = u.storage.Save(ctx, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

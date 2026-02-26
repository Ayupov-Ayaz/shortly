package shortener

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/ayupov-ayaz/shortly/internal/entity"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
	"github.com/ayupov-ayaz/shortly/internal/storage"
)

type Shortener interface {
	ShortenURL(
		ctx context.Context,
		req *entity.ShortURLRequest,
	) (*entity.URLResponse, error)
}

type URLShortener struct {
	storage   storage.Storage
	generator id.Generator

	baseURL       *url.URL //todo: use baseURL
	defaultExpire time.Duration
}

func New(
	storage storage.Storage,
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
	req *entity.ShortURLRequest,
) (*entity.URLResponse, error) {
	// if original url already exists, return it
	resp, err := u.storage.GetByOrigin(ctx, req.URL)
	if err != nil && !errors.Is(err, storage.ErrNotFound) {
		return nil, err
	} else if resp != nil {
		return resp, nil
	}

	shortURL := u.baseURL.JoinPath(u.generator.Generate())

	now := time.Now().UTC()
	expireAt := u.calculateExpireAt(now, req.Expire)

	res := entity.NewURLResponse(
		req.URL,
		shortURL.String(),
		now,
		expireAt,
	)

	err = u.storage.Save(ctx, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *URLShortener) calculateExpireAt(
	now time.Time,
	expire time.Duration,
) time.Time {
	if expire == 0 {
		expire = u.defaultExpire
	}

	return now.Add(expire)
}

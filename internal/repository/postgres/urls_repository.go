package postgres

import (
	"context"
	"fmt"

	"github.com/ayupov-ayaz/shortly/internal/entity"
)

type URLsRepository struct {
	pgx *ConnPool
}

const active = true

func NewURLsRepository(pgx *ConnPool) *URLsRepository {
	return &URLsRepository{pgx: pgx}
}

func (r *URLsRepository) Create(
	ctx context.Context, shortURL entity.URL,
) error {
	const query = `
	INSERT INTO urls (original_url, short_code, expires_at, created_at, active)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.pgx.Exec(ctx, query,
		shortURL.OriginalURL,
		shortURL.ShortCode,
		shortURL.ExpiresAt,
		shortURL.CreatedAt,
		active)
	if err != nil {
		return fmt.Errorf("pgx.Exec: %w", err)
	}

	return nil
}

func (r *URLsRepository) GetByOrigin(
	ctx context.Context, shortURL string,
) (*entity.URL, error) {
	return nil, nil
}

func (r *URLsRepository) GetByShortURL(
	ctx context.Context, originalURL string,
) (*entity.URL, error) {
	return nil, nil
}

func (r *URLsRepository) Delete(
	ctx context.Context, shortURL string,
) error {
	return nil
}

func (r *URLsRepository) Update(
	ctx context.Context, req entity.URL,
) error {
	return nil
}

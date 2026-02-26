package storage

import (
	"context"

	"github.com/ayupov-ayaz/shortly/internal/entity"
)

type StubStorage struct{}

func NewStubStorage() *StubStorage {
	return &StubStorage{}
}

func (s *StubStorage) Save(
	ctx context.Context,
	req *entity.URLResponse,
) error {
	return nil
}

func (s *StubStorage) GetByOrigin(
	ctx context.Context,
	shortURL string,
) (*entity.URLResponse, error) {
	return nil, nil
}

func (s *StubStorage) GetByShortURL(
	ctx context.Context,
	originalURL string,
) (*entity.URLResponse, error) {
	return nil, nil
}

func (s *StubStorage) Del(
	ctx context.Context,
	shortURL string,
) error {
	return nil
}

package repository

import (
	"context"
	"sync"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
)

type InMemoryRepository struct {
	mu               sync.Mutex
	shortURLStore    map[string]*gen.CreateURLResponse
	originToShortURL map[string]*gen.CreateURLResponse
}

func NewStubRepository() *InMemoryRepository {
	return &InMemoryRepository{
		shortURLStore:    make(map[string]*gen.CreateURLResponse),
		originToShortURL: make(map[string]*gen.CreateURLResponse),
	}
}

func (r *InMemoryRepository) Save(
	ctx context.Context,
	resp *gen.CreateURLResponse,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shortURLStore[resp.ShortURL] = resp
	r.originToShortURL[resp.OriginalURL] = resp

	return nil
}

func (r *InMemoryRepository) GetByOrigin(
	ctx context.Context,
	originURL string,
) (*gen.CreateURLResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp, ok := r.originToShortURL[originURL]
	if !ok {
		return nil, ErrNotFound
	}

	return resp, nil
}

func (r *InMemoryRepository) GetByShortURL(
	ctx context.Context,
	shortURL string,
) (*gen.CreateURLResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp, ok := r.shortURLStore[shortURL]
	if !ok {
		return nil, ErrNotFound
	}

	return resp, nil
}

func (r *InMemoryRepository) Del(
	ctx context.Context,
	shortURL string,
) error {
	r.mu.Lock()
	r.mu.Unlock()

	resp, ok := r.shortURLStore[shortURL]
	if !ok {
		return nil
	}

	delete(r.shortURLStore, shortURL)
	delete(r.originToShortURL, resp.OriginalURL)

	return nil
}

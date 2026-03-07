package stub

import (
	"context"
	"sync"

	"github.com/ayupov-ayaz/shortly/internal/api/gen"
	"github.com/ayupov-ayaz/shortly/internal/repository"
)

type URLInMemoryRepository struct {
	mu               sync.Mutex
	shortURLStore    map[string]*gen.CreateURLResponse
	originToShortURL map[string]*gen.CreateURLResponse
}

func NewStubRepository() *URLInMemoryRepository {
	return &URLInMemoryRepository{
		shortURLStore:    make(map[string]*gen.CreateURLResponse),
		originToShortURL: make(map[string]*gen.CreateURLResponse),
	}
}

func (r *URLInMemoryRepository) Save(
	ctx context.Context,
	resp *gen.CreateURLResponse,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shortURLStore[resp.ShortURL] = resp
	r.originToShortURL[resp.OriginalURL] = resp

	return nil
}

func (r *URLInMemoryRepository) GetByOrigin(
	ctx context.Context,
	originURL string,
) (*gen.CreateURLResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp, ok := r.originToShortURL[originURL]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return resp, nil
}

func (r *URLInMemoryRepository) GetByShortURL(
	ctx context.Context,
	shortURL string,
) (*gen.CreateURLResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp, ok := r.shortURLStore[shortURL]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return resp, nil
}

func (r *URLInMemoryRepository) Del(
	ctx context.Context,
	shortURL string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	resp, ok := r.shortURLStore[shortURL]
	if !ok {
		return nil
	}

	delete(r.shortURLStore, shortURL)
	delete(r.originToShortURL, resp.OriginalURL)

	return nil
}

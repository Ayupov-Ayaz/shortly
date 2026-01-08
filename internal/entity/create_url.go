package entity

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type CreateURLRequest struct {
	URL       string        `json:"url"`
	CustomKey string        `json:"customKey,omitempty"`
	Expire    time.Duration `json:"expiredAt,omitempty"`
}

type URLResponse struct {
	OriginalURL string    `json:"originalURL"`
	ShortURL    string    `json:"shortURL"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func (c *CreateURLRequest) Validate() error {
	if strings.TrimSpace(c.URL) == "" {
		return ErrURLIsEmpty
	}

	_, err := url.Parse(c.URL)
	if err != nil {
		return fmt.Errorf("parsing original url: %w", err)
	}

	return nil
}

func NewCreateURLResponse(
	originalURL, shortURL string,
	createdAt, expiresAt time.Time,
) *URLResponse {
	return &URLResponse{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		CreatedAt:   createdAt,
		ExpiresAt:   expiresAt,
	}
}

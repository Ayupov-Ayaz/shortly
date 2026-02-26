package api

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/ayupov-ayaz/shortly/internal/api/rest"
	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
	"github.com/ayupov-ayaz/shortly/internal/storage"
)

func Configure(
	router fiber.Router,
	cfg *config.Config,
) error {
	generator, err := id.NewSnowflakeGenerator(cfg.APP.Shortener.NodeID)
	if err != nil {
		return fmt.Errorf("creating snowflake generator: %w", err)
	}

	baseURL, err := cfg.APP.BaseURL()
	if err != nil {
		return fmt.Errorf("parsing base url: %w", err)
	}

	srv := shortener.New(storage.NewStubStorage(),
		generator,
		baseURL,
		cfg.APP.URLShortenerTTL(),
	)

	rest.NewURLShortenerHandler(srv).RegisterRouter(router)

	return nil
}

package api

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/repository/stub"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest/handler"
)

func Configure(
	router chi.Router,
	cfg *config.Config,
) error {
	respWriter := handler.NewResponseWriter()

	generator, err := id.NewSnowflakeGenerator(cfg.APP.Shortener.NodeID)
	if err != nil {
		return fmt.Errorf("creating snowflake generator: %w", err)
	}

	baseURL, err := cfg.Server.BaseURL(cfg.Env)
	if err != nil {
		return fmt.Errorf("parsing base url: %w", err)
	}

	shortenerSrv := shortener.New(stub.NewInMemoryRepository(),
		generator,
		baseURL,
		cfg.APP.ShortURLsTTL(),
	)

	configureShortener(respWriter, router, shortenerSrv)
	configureLiveness(respWriter, router)
	configureSwagger(respWriter, router)

	return nil
}

package api

import (
	"context"
	"fmt"
	"time"

	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/helper/environment"
	"github.com/ayupov-ayaz/shortly/internal/repository/postgres"
	"github.com/ayupov-ayaz/shortly/internal/service/id"
	"github.com/ayupov-ayaz/shortly/internal/service/shortener"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest/handler"
)

func Configure(
	server *Server,
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

	pgxPool, err := configurePostgres(cfg.Env, server, cfg.Postgres)
	if err != nil {
		return err
	}

	server.SetPostgresCloser(pgxPool)

	urlsRepository := postgres.NewURLsRepository(pgxPool)
	now := func() time.Time {
		return time.Now().UTC()
	}

	shortenerSrv := shortener.New(
		urlsRepository,
		generator,
		baseURL,
		now,
	)

	configureShortener(respWriter, server.router, shortenerSrv)
	configureLiveness(respWriter, server.router)
	configureSwagger(respWriter, server.router)

	return nil
}

func configurePostgres(
	env environment.Env, server *Server, cfg config.Postgres,
) (*postgres.ConnPool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(), cfg.ConnectTimeout)
	defer cancel()

	pgxPool, err := postgres.NewConnPool(ctx, postgres.Config{
		Env:               env,
		DSN:               cfg.DSN(),
		MaxConns:          cfg.MaxConns,
		MinConns:          cfg.MinConns,
		MaxConnLifetime:   cfg.MaxConnLifetime,
		MaxConnIdleTime:   cfg.MaxConnIdleTime,
		HealthCheckPeriod: cfg.HealthCheckPeriod,
	})

	if err != nil {
		return nil, fmt.Errorf("connection to postgres: %w", err)
	}

	return pgxPool, nil
}

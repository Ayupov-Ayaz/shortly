package rest

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/ayupov-ayaz/shortly/internal/config"
)

func setupCors(
	env string,
	domains []string,
) fiber.Handler {
	cfg := cors.Config{
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodOptions,
		},
		AllowHeaders: []string{
			fiber.HeaderOrigin,
			fiber.HeaderContentType,
			fiber.HeaderAccept,
			fiber.HeaderAuthorization,
		},
		AllowOrigins: domains,
	}

	const (
		devMaxAge  = 60 * 60      // 1h
		prodMaxAge = 60 * 60 * 24 // 24h
	)

	switch env {
	case config.EnvDevelopment:
		cfg.MaxAge = devMaxAge
		cfg.AllowCredentials = false // for use "*"
	case config.EnvProduction:
		cfg.MaxAge = prodMaxAge
		cfg.AllowCredentials = true
	}

	return cors.New(cfg)
}

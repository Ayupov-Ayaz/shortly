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
	}

	switch env {
	case config.EnvDevelopment:
		cfg.MaxAge = 60 * 60         // 1h
		cfg.AllowCredentials = false // for use "*"
		cfg.AllowOrigins = []string{"*"}
	case config.EnvProduction:
		cfg.AllowOrigins = domains
		cfg.MaxAge = 60 * 60 * 24 // 24h
		cfg.AllowCredentials = true
	}

	return cors.New(cfg)
}

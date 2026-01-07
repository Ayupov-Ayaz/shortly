package rest

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type Config struct {
	AppName string
	Env     string
	Domains []string
}

func New(
	cfg Config,
) (*fiber.App, error) {
	logFile, err := panicLogFile(cfg.Env)
	if err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
		AppName:      cfg.AppName,
		ErrorHandler: errorHandler,
	})

	_ = app.Use(
		setupRecovery(logFile),
		setupCors(cfg.Env, cfg.Domains),
	)

	return app, nil
}

// todo: implement graceful shutdown
// todo: logger middleware

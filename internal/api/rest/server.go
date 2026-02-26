package rest

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Config struct {
	PanicLogFilePath string

	AppName string
	Env     string
	Domains []string
}

func New(
	cfg Config,
) (*fiber.App, error) {
	panicLogFile, err := createPanicLogFile(cfg.PanicLogFilePath)
	if err != nil {
		return nil, fmt.Errorf("creating log file for panics: %w", err)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
		AppName:      cfg.AppName,
		ErrorHandler: errorHandler,
	})

	app.Use(
		setupRecovery(panicLogFile),
		setupCors(cfg.Env, cfg.Domains),
	)

	return app, nil
}

// todo: implement graceful shutdown
// todo: logger middleware

package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/ayupov-ayaz/shortly/internal/config"
)

type Config struct {
	Env          string
	LivenessPath string
	Timeout      time.Duration
}

func NewServer(cfg Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(cfg.Timeout),
		middleware.Heartbeat(cfg.LivenessPath),
		cors.Handler(corsOptions(cfg.Env)),
	)

	return router
}

func corsOptions(env string) cors.Options {
	allowOrigins := []string{"https://*"}

	if env == config.EnvDevelopment {
		allowOrigins = append(allowOrigins, "http://*")
	}

	return cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: []string{
			"GET", "POST", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type",
		},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		// todo: do i need to configure specific ttl for prod?
		MaxAge: 60 * 5, // 5 min in seconds
	}
}

package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouterConfig struct {
	AllowOrigins []string
	Timeout      time.Duration
}

func NewRouter(cfg RouterConfig) *chi.Mux {
	resp := chi.NewRouter()

	resp.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(cfg.Timeout),
		middleware.AllowContentType("application/json"),
		cors.Handler(corsOptions(cfg.AllowOrigins)),
	)

	return resp
}

func corsOptions(allowOrigins []string) cors.Options {
	return cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: []string{
			"GET", "POST", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type",
		},
		AllowCredentials: true,
		// todo: do i need to configure specific ttl for prod?
		MaxAge: 60 * 5, // 5 min in seconds
	}
}

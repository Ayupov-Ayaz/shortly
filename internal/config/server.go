package config

import (
	"fmt"
	"net/url"
)

type Server struct {
	Host         string   `env:"HOST" envDefault:"localhost"`
	AllowOrigins []string `env:"ALLOW_ORIGINS" validate:"required,min=1"`
	Port         int      `env:"PORT" envDefault:"9000"`
	Timeout      Timeout  `envPrefix:"TIMEOUT_"`
}

func (cfg Server) BaseURL(env string) (*url.URL, error) {
	var scheme string

	switch env {
	case EnvDevelopment:
		scheme = "http"
	case EnvProduction:
		scheme = "https"
	default:
		return nil, fmt.Errorf("invalid environment: %s", env)
	}

	baseURL, err := url.Parse(
		fmt.Sprintf("%s://%s", scheme, cfg.ListenAddr()))
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %w", err)
	}

	return baseURL, nil
}

func (cfg Server) ListenAddr() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

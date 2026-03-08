package config

import (
	"fmt"
	"net/url"

	"github.com/ayupov-ayaz/shortly/internal/helper/environment"
)

type Server struct {
	Host         string   `env:"HOST" envDefault:"localhost"`
	AllowOrigins []string `env:"ALLOW_ORIGINS" validate:"required,min=1"`
	Port         int      `env:"PORT" envDefault:"9000"`
	Timeout      Timeout  `envPrefix:"TIMEOUT_"`
}

func (cfg Server) BaseURL(env environment.Env) (*url.URL, error) {
	scheme := "https"

	if env.IsDevelopment() {
		scheme = "http"
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

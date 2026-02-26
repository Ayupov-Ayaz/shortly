package config

import (
	"fmt"
	"net/url"
	"time"
)

type APP struct {
	CORS      CORS      `envPrefix:"CORS_"`
	Shortener Shortener `envPrefix:"SHORTENER_"`

	Env       string `env:"ENV" envDefault:"development"`
	Host      string `env:"HOST" envDefault:"localhost"`
	PanicFile string `env:"PANIC_FILE" envDefault:"logs/panic.log"`
	Port      int    `env:"PORT" envDefault:"9000"`
}

type CORS struct {
	Domains []string `env:"DOMAINS" envDefault:"" envSeparator:","`
}

type Shortener struct {
	NodeID int64 `env:"NODE_ID" envDefault:"1"`
	TTL    int64 `env:"URL_TTL" envDefault:"86400"` // 24h in seconds
}

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func (a APP) BaseURL() (*url.URL, error) {
	var scheme string

	switch a.Env {
	case EnvDevelopment:
		scheme = "http"
	case EnvProduction:
		scheme = "https"
	default:
		return nil, fmt.Errorf("invalid environment: %s", a.Env)
	}

	baseURL, err := url.Parse(
		fmt.Sprintf("%s://%s", scheme, a.ListenAddr()))
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %w", err)
	}

	return baseURL, nil
}

func (a APP) ListenAddr() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func (a APP) URLShortenerTTL() time.Duration {
	return time.Duration(a.Shortener.TTL) * time.Second
}

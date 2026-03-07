package config

import (
	"time"
)

type APP struct {
	Shortener Shortener `envPrefix:"SHORTENER_"`
	// timeout for business logic
	Timeout time.Duration `env:"TIMEOUT" envDefault:"4s"`
	// todo: analytics
	// todo: redirect
}

type Shortener struct {
	NodeID int64 `env:"NODE_ID" envDefault:"1"`
	TTL    int64 `env:"URL_TTL" envDefault:"86400"` // 24h in seconds
}

func (cfg APP) ShortURLsTTL() time.Duration {
	return time.Duration(cfg.Shortener.TTL) * time.Second
}

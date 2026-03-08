package config

import (
	"fmt"
	"time"
)

type Postgres struct {
	Host     string `env:"HOST" envDefault:"shortly"`
	User     string `env:"USER" validate:"required"`
	Password string `env:"PASSWORD" validate:"required"`
	Name     string `env:"NAME" envDefault:"shortly"`
	SSLMode  string `env:"SSLMODE" envDefault:"disabled"`
	Port     int    `env:"PORT" envDefault:"5432"`

	// Pool settings
	MaxConns          int32         `env:"MAX_CONNS" envDefault:"25"`
	MinConns          int32         `env:"MIN_CONNS" envDefault:"5"`
	MaxConnLifetime   time.Duration `env:"MAX_CONN_LIFETIME" envDefault:"1h"`
	MaxConnIdleTime   time.Duration `env:"MAX_CONN_IDLE_TIME" envDefault:"30m"`
	HealthCheckPeriod time.Duration `env:"HEALTH_CHECK_PERIOD" envDefault:"1m"`
	ConnectTimeout    time.Duration `env:"CONNECT_TIMEOUT" envDefault:"5s"`
}

func (p Postgres) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.Name, p.SSLMode,
	)
}

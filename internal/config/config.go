package config

import "github.com/ayupov-ayaz/shortly/internal/helper/environment"

type Config struct {
	Env environment.Env `env:"ENV" envDefault:"development"`

	APP      APP      `envPrefix:"APP_"`
	Server   Server   `envPrefix:"SERVER_"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Redis    Redis    `envPrefix:"REDIS_"`
}

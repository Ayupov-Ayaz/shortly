package config

type Config struct {
	Env string `env:"ENV" envDefault:"development"`

	APP      APP      `envPrefix:"APP_"`
	Server   Server   `nevPrefix:"SERVER"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Redis    Redis    `envPrefix:"REDIS_"`
}

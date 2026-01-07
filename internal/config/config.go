package config

type Config struct {
	APP      APP      `envPrefix:"APP_"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Redis    Redis    `envPrefix:"REDIS_"`
}


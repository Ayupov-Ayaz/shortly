package config

type Config struct {
	APP      APP      `envPrefix:"APP_"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Redis    Redis    `envPrefix:"REDIS_"`
}

type APP struct {
	Env  string `env:"ENV" envDefault:"development"`
	Port string `env:"PORT" envDefault:"9000"`
}

type Postgres struct {
	Host     string `env:"HOST" envDefault:"shortly"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME" envDefault:"shortly"`
	SSLMode  string `env:"SSLMODE" envDefault:"disabled"`
}

type Redis struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"6379"`
	Password string `env:"PASSWORD" envDefault:""`
	DB       int    `env:"DB" envDefault:"0"`
}

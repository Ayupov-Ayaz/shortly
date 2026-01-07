package config

type Postgres struct {
	Host     string `env:"HOST" envDefault:"shortly"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME" envDefault:"shortly"`
	SSLMode  string `env:"SSLMODE" envDefault:"disabled"`
}

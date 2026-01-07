package config

import "strings"

type APP struct {
	Env     string   `env:"ENV" envDefault:"development"`
	Port    string   `env:"PORT" envDefault:"9000"`
	Domains []string `env:"DOMAINS" envDefault:"" envSeparator:","`
}

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func (a APP) ServerPort() string {
	const prefix = ":"

	if strings.HasPrefix(a.Port, prefix) {
		return a.Port
	}

	return prefix + a.Port
}

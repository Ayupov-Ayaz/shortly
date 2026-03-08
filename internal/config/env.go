package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func FromEnv(fileName, prefix string) (*Config, error) {
	err := loadDotEnvFile(fileName)
	if err != nil {
		return nil, err
	}

	cfg, err := parseEnv(prefix)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadDotEnvFile(fileName string) error {
	if err := godotenv.Load(fileName); err != nil {
		return fmt.Errorf("loading '%s' file: %w", fileName, err)
	}

	return nil
}

func parseEnv(prefix string) (*Config, error) {
	opt := env.Options{
		RequiredIfNoDef: true,
		Prefix:          prefix,
	}

	cfg := new(Config)

	if err := env.ParseWithOptions(cfg, opt); err != nil {
		return nil, fmt.Errorf("parsing env: %w", err)
	}

	return cfg, nil
}

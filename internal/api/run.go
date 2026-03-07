package api

import (
	"flag"
	"fmt"

	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/transport/rest"
)

const (
	prefix = "SHORTLY_"
)

func Run() error {
	cfg, err := readConfig()
	if err != nil {
		return err
	}

	router := rest.NewRouter(rest.RouterConfig{
		AllowOrigins: cfg.Server.AllowOrigins,
		Timeout:      cfg.APP.Timeout,
	})

	server := NewServer(router, ServerConfig{
		Addr:         cfg.Server.ListenAddr(),
		ReadTimeout:  cfg.Server.Timeout.Read,
		WriteTimeout: cfg.Server.Timeout.Write,
		IdleTimeout:  cfg.Server.Timeout.Idle,
	})

	err = Configure(router, cfg)
	if err != nil {
		return fmt.Errorf("configuring api: %w", err)
	}

	return server.Start()
}

func readConfig() (*config.Config, error) {
	env := parseEnvFlag()
	envFileName := choseEnvFile(env)

	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		return nil, fmt.Errorf("parsing env: %w", err)
	}

	return cfg, nil
}

func parseEnvFlag() string {
	env := flag.String("env", "development", "environment")
	flag.Parse()

	return *env
}

func choseEnvFile(env string) string {
	const (
		devEnvFile  = ".env.development"
		prodEnvFile = ".env.production"
	)

	switch env {
	case config.EnvDevelopment:
		return devEnvFile
	default:
		return prodEnvFile
	}
}

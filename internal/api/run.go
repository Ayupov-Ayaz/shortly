package api

import (
	"flag"
	"fmt"

	"github.com/ayupov-ayaz/shortly/internal/config"
	"github.com/ayupov-ayaz/shortly/internal/helper/environment"
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

	err = Configure(server, cfg)
	if err != nil {
		return fmt.Errorf("configuring api: %w", err)
	}

	return server.Start()
}

func readConfig() (*config.Config, error) {
	env := parseEnvFlag()
	envFileName, err := choseEnvFile(env)
	if err != nil {
		return nil, err
	}

	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		return nil, fmt.Errorf("parsing env: %w", err)
	}

	cfg.Env = env

	return cfg, nil
}

func parseEnvFlag() environment.Env {
	env := flag.String("env", "development", "environment")
	flag.Parse()

	return environment.Env(*env)
}

func choseEnvFile(env environment.Env) (string, error) {
	const (
		devEnvFile  = ".env.development"
		prodEnvFile = ".env.production"
	)

	if env.IsDevelopment() {
		return devEnvFile, nil
	}

	if env.IsProduction() {
		return prodEnvFile, nil
	}

	return "", fmt.Errorf("unknown environment: %s", env)
}

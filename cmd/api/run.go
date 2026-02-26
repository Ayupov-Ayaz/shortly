package api

import (
	"flag"
	"fmt"

	"github.com/ayupov-ayaz/shortly/internal/api"
	"github.com/ayupov-ayaz/shortly/internal/api/rest"
	"github.com/ayupov-ayaz/shortly/internal/config"
)

const (
	appName = "shortly"
	prefix  = "SHORTLY_"
)

func Run() error {
	env := parseEnvFlag()
	envFileName := choseEnvFile(env)

	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		return fmt.Errorf("parsing env: %w", err)
	}

	app, err := rest.New(rest.Config{
		AppName:          appName + ":" + env,
		Env:              env,
		Domains:          cfg.APP.CORS.Domains,
		PanicLogFilePath: cfg.APP.PanicFile,
	})
	if err != nil {
		return fmt.Errorf("creating rest server: %w", err)
	}

	err = api.Configure(app, cfg)
	if err != nil {
		return fmt.Errorf("configuring api: %w", err)
	}

	// todo: add graceful shutdown
	if err = app.Listen(cfg.APP.ListenAddr()); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}

	return nil
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

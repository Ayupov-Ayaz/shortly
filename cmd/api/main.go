package main

import (
	"flag"

	"github.com/ayupov-ayaz/shortly/internal/api/rest"
	"github.com/ayupov-ayaz/shortly/internal/config"
)

const (
	appName = "shortly"
	prefix  = "SHORTLY_"
)

func main() {
	env := parseEnvFlag()
	envFileName := choseEnvFile(env)

	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		panic(err)
	}

	srv, err := rest.New(rest.Config{
		AppName: appName + ":" + env,
		Env:     env,
		Domains: cfg.APP.Domains,
	})
	if err != nil {
		panic(err)
	}

	if err = srv.Listen(cfg.APP.ServerPort()); err != nil {
		panic(err)
	}
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

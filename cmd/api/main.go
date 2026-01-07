package main

import (
	"fmt"

	"github.com/ayupov-ayaz/shortly/internal/config"
)

const (
	prefix      = "SHORTLY_"
	envFileName = ".env.development"
)

func main() {
	cfg, err := config.FromEnv(envFileName, prefix)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}

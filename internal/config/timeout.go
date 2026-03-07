package config

import "time"

type Timeout struct {
	Read  time.Duration `env:"READ" envDefault:"5s" validate:"min=1s"`
	Write time.Duration `env:"WRITE" envDefault:"5s" validate:"min=1s"`
	Idle  time.Duration `env:"IDLE" envDefault:"60s" validate:"min=30s"`
}

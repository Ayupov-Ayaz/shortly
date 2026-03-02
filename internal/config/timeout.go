package config

import "time"

type Timeout struct {
	Read  time.Duration `env:"READ" envDefault:"5s"`
	Write time.Duration `env:"WRITE" envDefault:"5s"`
	Idle  time.Duration `env:"IDLE" envDefault:"60s"`
}

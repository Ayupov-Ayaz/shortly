package main

import (
	"github.com/ayupov-ayaz/shortly/cmd/api"
)

func main() {
	if err := api.Run(); err != nil {
		panic(err)
	}
}

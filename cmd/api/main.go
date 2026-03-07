package main

import (
	"log"

	"github.com/ayupov-ayaz/shortly/internal/api"
)

func main() {
	if err := api.Run(); err != nil {
		log.Fatalf("❌ Application failed: %v", err)
	}
}

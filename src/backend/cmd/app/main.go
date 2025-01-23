package main

import (
	"backend/config"
	"backend/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}
	cfg.App.Name = "My App"
	app.Run(cfg)
}

package main

import (
	"github.com/sw1pr0g/sw1pr0g-website/config"
	"github.com/sw1pr0g/sw1pr0g-website/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}

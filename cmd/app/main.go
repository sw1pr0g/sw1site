package main

import (
	"github.com/sw1pr0g/sw1site/config"
	"github.com/sw1pr0g/sw1site/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}

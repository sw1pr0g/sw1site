package app

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1site/config"
	"github.com/sw1pr0g/sw1site/internal/components"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {
	app.Route("/", &components.Home{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "sw1pr0g",
		Description: "sw1pr0g website",
		Icon: app.Icon{
			Default: "static/icon.png", // Specify default favicon.
		},
	})

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}

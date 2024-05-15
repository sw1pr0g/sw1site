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
	app.Route("/profile", &components.Profile{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "sw1pr0g",
		Description: "sw1pr0g website",
		Icon: app.Icon{
			Default: "static/icon.png", // Specify default favicon.
		},
		LoadingLabel: "Lofi music player to work, study and relax.",
		Image:        "https://lofimusic.app/web/covers/lofimusic.png",
	})

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}

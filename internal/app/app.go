package app

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1site/config"
	"github.com/sw1pr0g/sw1site/internal/components"
	"log"
	"mime"
	"net/http"
)

func Run(cfg *config.Config) {
	app.Route("/", &components.Base{})
	app.Route("/home", &components.Home{})
	app.Route("/blog", &components.Blog{})
	app.RunWhenOnBrowser()

	if err := mime.AddExtensionType(".wasm", "application/wasm"); err != nil {
		log.Fatal(err)
	}

	staticDir := "/static/web"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/web/", http.StripPrefix("/static/web/", fs))

	http.Handle("/", &app.Handler{
		Name: cfg.App.Name,
	})

	log.Printf("Starting server on http://localhost:%s", cfg.HTTP.Port)
	if err := http.ListenAndServe(":"+cfg.HTTP.Port, nil); err != nil {
		log.Fatal(err)
	}
}

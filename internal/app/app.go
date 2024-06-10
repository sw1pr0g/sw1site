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
	app.Route("/", &App{})

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

type App struct {
	app.Compo
}

func (a *App) Render() app.UI {
	return app.Div().Body(
		&components.Header{},
		app.Div().Body(
			app.If(
				app.Window().URL().Path == "/",
				&components.Home{},
			).ElseIf(
				app.Window().URL().Path == "/profile",
				&components.Profile{},
			),
		),
	)
}

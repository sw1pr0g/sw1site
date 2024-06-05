package app

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/sw1pr0g/sw1site/config"
	"log"
	"net/http"
)

// Run запускает приложение с заданной конфигурацией
func Run(cfg *config.Config) {
	// Регистрация маршрутов приложения
	app.Route("/", &home{})
	app.Route("/profile", &profile{})
	app.RunWhenOnBrowser()

	// Настройка сервера для обслуживания статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Настройка основного приложения
	http.Handle("/", &app.Handler{
		Name:        cfg.App.Name,
		Description: "An example",
	})

	log.Printf("Starting server on :%s", cfg.HTTP.Port)
	if err := http.ListenAndServe(":"+cfg.HTTP.Port, nil); err != nil {
		log.Fatal(err)
	}
}

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Hello, World!"),
		app.P().Text("This is a home page."),
	)
}

type profile struct {
	app.Compo
}

func (p *profile) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Profile Page"),
		app.P().Text("This is a profile page."),
	)
}

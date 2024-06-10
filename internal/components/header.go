package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Header().Body(
		app.Div().Class("logo").Body(
			app.H1().Text("Alex Poryadin"),
		),
		app.Nav().Body(
			app.Ul().Body(
				app.Li().Body(
					app.A().Href("/blog").Text("Blog"),
				),
				app.Li().Body(
					app.A().Href("/newsletter").Text("Newsletter"),
				),
				app.Li().Body(
					app.A().Href("/contacts").Text("Contacts"),
				),
			),
		),
	)
}

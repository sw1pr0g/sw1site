package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.H1().
			Class("title").
			Text("Alex Poryadin"),
		app.P().
			Class("text").
			Text("Software Engineer"),
		app.Br(),
	)
}

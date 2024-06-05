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
			Text("Build a GUI with Go"),
		app.P().
			Class("text").
			Text("Just because Go and this package are really awesome!"),
	)
}

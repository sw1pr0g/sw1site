package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text("sw1prog's homepage"),
		app.P().
			Text("maybe you were waiting for the design, but I don't care about that :)"),
		app.Br(),
	)
}

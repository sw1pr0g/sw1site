package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Blog struct {
	app.Compo
}

func (h *Blog) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text("blog page"),
	)
}

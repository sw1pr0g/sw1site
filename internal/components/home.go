package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Text("Home Page")
}

package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Profile struct {
	app.Compo
}

func (p *Profile) Render() app.UI {
	return app.Div().Body(
		app.H2().Text("Profile"),
		app.P().Text("This is the profile page content."),
	)
}

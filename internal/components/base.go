package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Base struct {
	app.Compo
}

func (h *Base) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text("sw1prog's home.. maybe you were waiting for the design, but I don't care about that"),
		app.P().Body(
			app.Text("i have a lazy "),
			app.A().
				Text("blog").
				Href("/blog"),
			app.Text("and "),
			app.A().
				Text("telegram channel").
				Href("https://t.me/sw1pr0g_channel").
				Target("_blank"),
		),
		app.P().Body(
			app.Text("check out my "),
			app.A().
				Text("github").
				Href("https://github.com/sw1pr0g").
				Target("_blank"),
		),
		app.P().Body(
			app.Text("heh, "),
			app.A().
				Text("let's go to readable pages").
				Href("/home"),
			app.Text(", i like jokes :)"),
		),
	)
}

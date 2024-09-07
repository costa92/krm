package app

import "github.com/costa92/krm/pkg/app"

const commandDesc = "Launch a krm tools server"

func NewApp(name string) *app.App {
	application := app.NewApp(name, "krm-dev tools",
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
	)
	return application
}

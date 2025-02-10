package app

import "github.com/urfave/cli"

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "Minha primeira CLI"
	app.Version = "1.0.0"
	return app
}
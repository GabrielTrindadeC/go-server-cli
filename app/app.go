package app

import (
	"go-cli/app/cmd"

	"github.com/urfave/cli"
)

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "Minha primeira CLI"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de servidores na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: cmd.SearchIp,
		},
	}
	return app
}


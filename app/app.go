package app

import (
	"go-server-cli/app/cmd"

	"github.com/urfave/cli"
)

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "Minha primeira CLI"
	app.Version = "1.0.0"
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		}}
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de servidores na internet",
			Flags: flags,
			Action: cmd.SearchIp,
		},
		{
			Name: "server",
			Usage: "Busca o nome do servidor",
			Flags: flags,
			Action: cmd.SearchNs,
		},
		{
			Name: "ping",
			Usage: "Faz um ping em um IP específico",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "ip",
				},
			},
			Action: cmd.Ping,
		},
	}
	return app
}


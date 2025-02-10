package main

import (
	"go-server-cli/app"
	"os"
)
func main() {
	app := app.Init()
	if erro := app.Run(os.Args); erro != nil {
		panic(erro)
	}
}
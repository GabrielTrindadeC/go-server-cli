package main

import (
	"go-cli/app"
	"os"
)
func main() {
	app := app.Init()
	if erro := app.Run(os.Args); erro != nil {
		panic(erro)
	}
	
}
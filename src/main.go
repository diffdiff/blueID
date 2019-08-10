package main

import (
	"github.com/diffdiff/blueID/app"
	"github.com/diffdiff/blueID/app/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)

	app.Run(":3000")
}

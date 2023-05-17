package main

import (
	"os"

	"github.com/kirebyte/exiftags-go/controller"
	"github.com/kirebyte/exiftags-go/ui"
	"github.com/urfave/cli/v2"
)

func main() {
	//initialize controller
	app := ui.New(controller.NewExifActionFunc())

	//start app and log in case of failure
	if err := app.Run(os.Args); err != nil {
		cli.Exit(err.Error(), 1)
	}
}

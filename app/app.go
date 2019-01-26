package app

import (
	"github.com/codegangsta/cli"

	command "github.com/kcwebapply/spg/command"
	config "github.com/kcwebapply/spg/config"
)

var app *cli.App

func AppInit() *cli.App {
	config := config.GetConfig()

	app = cli.NewApp()
	app.Name = config.App.Name
	app.Usage = "github.com/cli/spg"
	app.Version = config.App.Version
	app.Commands = []cli.Command{
		{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "create spring-boot package from .toml file.",
			Action:  command.GeneratePackage,
		},
	}
	return app
}

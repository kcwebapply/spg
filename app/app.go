package app

import (
	"github.com/codegangsta/cli"

	command "github.com/kcwebapply/spg/command"
	config "github.com/kcwebapply/spg/config"
)

var app *cli.App

// AppInit initialize terminal setting.
func AppInit() *cli.App {
	config := config.GetConfig()

	app = cli.NewApp()
	app.Name = config.App.Name
	app.Usage = "github.com/cli/spg"
	app.Version = config.App.Version

	initFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "n,name",
			Value: "spring-sample",
			Usage: "name",
		},
		cli.StringFlag{
			Name:  "a,artifactId",
			Value: "spring-sample",
			Usage: "artifactId",
		},

		cli.StringFlag{
			Name:  "g,groupId",
			Value: "com.sample",
			Usage: "groupId",
		},
		cli.StringFlag{
			Name:  "s,spring",
			Value: "2.1.2.RELEASE",
			Usage: "springVersion",
		},
		cli.StringFlag{
			Name:  "j,java",
			Value: "1.8",
			Usage: "javaVersion",
		},

		cli.StringFlag{
			Name:  "jdbc,db",
			Value: "",
			Usage: "(add setting on application.properties and dependency) dbConnectionSetting",
		},

		cli.StringFlag{
			Name:  "table,entity",
			Value: "",
			Usage: "(generate class file) EntityName",
		},

		cli.StringFlag{
			Name:  "driver",
			Value: "",
			Usage: "(add setting on application.properties) driverSetting.",
		},

		cli.StringFlag{
			Name:  "schedule",
			Value: "",
			Usage: "(generating class file) add scheduler Bean Class.",
		},

		cli.StringFlag{
			Name:  "zone",
			Value: "",
			Usage: "(editing class file) add cron setting on `@Scheduler` annotation.",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "genf",
			Aliases: []string{"f"},
			Usage:   "create spring-boot package from .toml file.",
			Action:  command.GeneratePackage,
		},

		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "generate package",
			Action:  command.GeneratePackage,
			Flags:   initFlag,
		},

		{
			Name:    "touch",
			Aliases: []string{"t"},
			Usage:   "generate default .toml file.",
			Action:  command.InitTomlFile,
			Flags:   initFlag,
		},
	}

	return app
}

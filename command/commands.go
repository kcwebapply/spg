package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	generator "github.com/kcwebapply/spg/generator"
	parser "github.com/kcwebapply/spg/parser"
)

func GeneratePackage(c *cli.Context) {
	fileName := c.Args().First()
	userInput := parser.Parse(fileName)

	fmt.Println("obj ", userInput.App.Name, " ", userInput.Db.Jdbc)
	fmt.Println("app", userInput.Db.Jdbc)
	if &userInput.App != nil {
		generator.GenerateMain(userInput.App)
	}

}

func parseConfig() {

}

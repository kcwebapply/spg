package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	generator "github.com/kcwebapply/spg/generator"
	parser "github.com/kcwebapply/spg/parser"
)

func GeneratePackage(c *cli.Context) {
	fileName := c.Args().First()
	userInput := parser.Parse(fileName)

	fmt.Println("obj ", userInput.App.Name, " ", userInput.Db.Jdbc)
	fmt.Println("app", userInput.Db.Jdbc)
	if &userInput.App != nil && &userInput.App.Name != nil {
		generator.CreateDirectory(userInput.App.Name)
		generator.GeneratePom(userInput)
		generator.GenerateMain(userInput.App.Name)
	} else {
		fmt.Println("please define `name` property on tomlFile.")
		os.Exit(0)
	}
}

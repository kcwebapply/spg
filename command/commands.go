package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	generator "github.com/kcwebapply/spg/generator"
	parser "github.com/kcwebapply/spg/parser"
)

// GeneratePackage generate spring-boot package.
func GeneratePackage(c *cli.Context) {
	fileName := c.Args().First()
	if &fileName == nil {
		fmt.Println("please input .toml file name.")
		os.Exit(0)
	}
	userInput := parser.Parse(fileName)
	if userInput.App.Name == "" {
		fmt.Println("please define `name` property on tomlFile.")
		os.Exit(0)
	}
	generator.CreateDirectory(userInput.App.Name)
	generator.GeneratePom(userInput)
	generator.GenerateMain(userInput.App.Name)
	generator.GeneratePropertiesFile(userInput)

	if userInput.Db.Driver != "" {
		generator.GenerateDB(userInput)
	}

}

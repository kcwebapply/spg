package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	generator.CreateDirectory(userInput)
	generator.GeneratePom(userInput)
	generator.GenerateMain(userInput)
	generator.GenerateTest(userInput)
	generator.GeneratePropertiesFile(userInput)

	if userInput.Db.Driver != "" {
		generator.GenerateDB(userInput)
	}

	fmt.Printf("\x1b[1;32mGenerating package %s completed!\x1b[0m\n", userInput.App.Name)

}

// InitTomlFile generate toml file.
func InitTomlFile(c *cli.Context) {
	// generate file.
	fileName := "spg.toml"
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)

	// open format file.
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err,", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("err,", err)
	}
	writer.Write(b)
}

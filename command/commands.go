package command

import (
	"bufio"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"strings"

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

	if userInput.Task.Schedule != "" {
		generator.GenerateTask(userInput)
	}

	fmt.Printf("\x1b[1;32mGenerating package %s completed!\x1b[0m\n", userInput.App.Name)

}

// InitTomlFile generate toml file.
func InitTomlFile(c *cli.Context) {
	name := "sample"
	groupId := "com.sample"
	artifactId := "sample"
	springVersion := "2.1.2.RELEASE"
	javaVersion := "1.8"

	if c.String("n") != "" {
		name = "\"" + c.String("n") + "\""
	}

	if c.String("a") != "" {
		artifactId = "\"" + c.String("a") + "\""
	}

	if c.String("g") != "" {
		groupId = "\"" + c.String("g") + "\""
	}

	if c.String("s") != "" {
		springVersion = "\"" + c.String("s") + "\""
	}

	if c.String("j") != "" {
		javaVersion = "\"" + c.String("j") + "\""
	}

	// generate file.
	fileName := "spg.toml"
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	defer writer.Flush()

	// open format file.
	f, err := os.Open(build.Default.GOPATH + "/src/github.com/kcwebapply/spg/default.toml")
	if err != nil {
		fmt.Println("err,", err)
		os.Exit(0)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("err,", err)
		os.Exit(0)
	}
	content := string(b)
	content = strings.Replace(content, "${name}", name, -1)
	content = strings.Replace(content, "${artifactId}", artifactId, -1)
	content = strings.Replace(content, "${groupId}", groupId, -1)
	content = strings.Replace(content, "${springVersion}", springVersion, -1)
	content = strings.Replace(content, "${javaVersion}", javaVersion, -1)
	writer.WriteString(content)

	fmt.Printf("\x1b[1;32mGenerating spg.toml file completed!\x1b[0m\n")
}

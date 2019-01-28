package generator

import (
	"go/build"
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
)

// GenerateMain touch mainClass.
func GenerateMain(userInput parser.UserInput) {
	appName := userInput.App.Name
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/main.java"
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userInput), -1)
	writer := generateFile(appName + path + "/" + getPathformatFromUserInput(userInput) + "/" + strcase.ToCamel(appName) + ".java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

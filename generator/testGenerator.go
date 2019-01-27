package generator

import (
	"go/build"
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
)

// GenerateTest touch mainClass.
func GenerateTest(userinput parser.UserInput) {
	appName := userinput.App.Name
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/test/java/test.java"
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userinput), -1)
	writer := generateFile(appName + "/src/test/java/" + getPathformatFromUserInput(userinput) + "/" + strcase.ToCamel(appName) + "Tests.java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

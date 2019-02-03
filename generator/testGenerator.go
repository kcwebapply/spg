package generator

import (
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
	template "github.com/kcwebapply/spg/template"
)

// GenerateTest touch mainClass.
func GenerateTestClass(userinput parser.UserInput) {
	appName := userinput.App.Name

	content := template.TEST
	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userinput), -1)
	writer := generateFile(appName + "/src/test/java/" + getPathformatFromUserInput(userinput) + "/" + strcase.ToCamel(appName) + "Tests.java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

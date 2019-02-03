package generator

import (
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
	template "github.com/kcwebapply/spg/template"
)

// GenerateMain touch mainClass.
func GenerateMainClass(userInput parser.UserInput) {
	appName := userInput.App.Name

	content := template.MAIN
	imports, annotations := getImports(userInput)

	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userInput), -1)
	content = strings.Replace(content, "${imports}", imports, -1)
	content = strings.Replace(content, "${annotations}", annotations, -1)

	writer := generateFile(appName + path + "/" + getPathformatFromUserInput(userInput) + "/" + strcase.ToCamel(appName) + ".java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

func getImports(userInput parser.UserInput) (string, string) {
	var imports = ""
	var annotations = ""
	if userInput.Task.Schedule != "" {
		imports += "import org.springframework.scheduling.annotation.EnableScheduling;\n"
		annotations += "@EnableScheduling"
	}
	return imports, annotations
}

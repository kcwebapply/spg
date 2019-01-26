package generator

import (
	"go/build"
	"strings"

	"github.com/iancoleman/strcase"
)

const path = "/src/main/java"

// GenerateMain touch mainClass.
func GenerateMain(appName string) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/main.java"
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	writer := generateFile(appName + path + "/" + strcase.ToCamel(appName) + ".java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
	template "github.com/kcwebapply/spg/template"
)

// GenerateDB touch entity and repository file.
func GenerateDBClass(userInput parser.UserInput) {
	// do nothing if tableName not designated here.
	tableName := userInput.Db.Table
	if &tableName == nil || tableName == "" {
		fmt.Println("please input table name.")
		return
	}

	modelPath := userInput.App.Name + path + "/" + getPathformatFromUserInput(userInput) + "/model"
	err := os.Mkdir(modelPath, 0777)
	if err != nil {
		fmt.Println("can'd make directory `model`.")
		os.Exit(0)
	}

	entityClassName := strcase.ToCamel(userInput.Db.Table) + "Entity"
	jpaClassName := strcase.ToCamel(userInput.Db.Table) + "Repository"
	// create Entity Class
	writer := generateFile(modelPath + "/" + entityClassName + ".java")
	content := template.ENTITY
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userInput), -1)
	content = strings.Replace(content, "${name}", entityClassName, -1)
	content = strings.Replace(content, "${tableName}", "\""+userInput.Db.Table+"\"", -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
	// create Jpa ckass
	writer = generateFile(modelPath + "/" + jpaClassName + ".java")
	content = template.JPA
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userInput), -1)
	content = strings.Replace(content, "${name}", jpaClassName, -1)
	content = strings.Replace(content, "${entityFileName}", entityClassName, -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

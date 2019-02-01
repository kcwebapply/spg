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
func GenerateDB(userInput parser.UserInput) {
	tableName := userInput.Db.Table
	if &tableName == nil || tableName == "" {
		fmt.Println("please input table name.")
		os.Exit(0)
	}
	err := os.Mkdir(userInput.App.Name+path+"/model", 0777)
	if err != nil {
		fmt.Println("can'd make directory `model`.")
		os.Exit(0)
	}

	entityClassName := strcase.ToCamel(userInput.Db.Table) + "Entity"
	jpaClassName := strcase.ToCamel(userInput.Db.Table) + "Repository"
	// create Entity Class
	writer := generateFile(userInput.App.Name + path + "/model/" + entityClassName + ".java")
	content := template.ENTITY
	content = strings.Replace(content, "${name}", entityClassName, -1)
	content = strings.Replace(content, "${tableName}", "\""+userInput.Db.Table+"\"", -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
	// create Jpa ckass
	writer = generateFile(userInput.App.Name + path + "/model/" + jpaClassName + ".java")
	content = template.JPA
	content = strings.Replace(content, "${name}", jpaClassName, -1)
	content = strings.Replace(content, "${entityFileName}", entityClassName, -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

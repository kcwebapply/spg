package generator

import (
	"bufio"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	parser "github.com/kcwebapply/spg/parser"
)

const path = "/src/main/java"

func CreateDirectory(appName string) {
	err := os.MkdirAll(appName+path, 0777)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = os.Mkdir(appName+"/src/main/resources", 0777)
	if err != nil {
		fmt.Println("err:", err)
	}
	os.Chmod(appName, 0777)
}

func GeneratePom(userInput parser.UserInput) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/pom.xml"
	content := getFormatFileContent(fileName)
	writer := generateFile(userInput.App.Name + "/pom.xml")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

func GeneratePropertiesFile(appName string) {
	writer := generateFile(appName + "/src/main/resources/application.properties")
	defer writer.Flush()
}

func GenerateMain(appName string) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/main.java"
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${name}", strcase.ToCamel(appName), -1)
	writer := generateFile(appName + path + "/main.java")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

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

	entityFileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/model/entity.java"
	jpaFileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/model/jpa.java"

	entityClassName := strcase.ToCamel(userInput.Db.Table) + "Entity"
	jpaClassName := strcase.ToCamel(userInput.Db.Table) + "Repository"
	// create Entity Class
	writer := generateFile(userInput.App.Name + path + "/model/" + entityClassName + ".java")
	content := getFormatFileContent(entityFileName)
	content = strings.Replace(content, "${name}", entityClassName, -1)
	content = strings.Replace(content, "${tableName}", "\""+userInput.Db.Table+"\"", -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))

	writer = generateFile(userInput.App.Name + path + "/model/" + jpaClassName + ".java")
	content = getFormatFileContent(jpaFileName)
	content = strings.Replace(content, "${name}", jpaClassName, -1)
	content = strings.Replace(content, "${entityFileName}", entityClassName, -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))

}

func getFormatFileContent(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err,", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("err,", err)
	}
	return string(b)
}

func generateFile(fileName string) *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}

package generator

import (
	"bufio"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

const path = "/src/main/java"

func CreateDirectory(appName string) {
	err := os.MkdirAll(appName+path, 0777)
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

func GenerateMain(appName string) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/main.java"
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${name}", appName, -1)
	writer := generateFile(appName + path + "/main.java")
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

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

func GenerateMain(app parser.App) {
	fmt.Println(app)
	//
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/main.java"
	content := getFileString(fileName)
	content = strings.Replace(content, "${name}", app.Name, -1)

	writer := getFileCleanWriter(fileName)
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

func getFileString(fileName string) string {
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

func getFileCleanWriter(fileName string) *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}

package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

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
func getPackageformatFromUserInput(userInput parser.UserInput) string {
	packagePath := userInput.App.GroupId + "." + userInput.App.ArtifactId
	return strings.Replace(packagePath, "-", "", -1)
}

func getPathformatFromUserInput(userInput parser.UserInput) string {
	groupIDPath := strings.Replace(userInput.App.GroupId, ".", "/", -1)
	groupIDPath = strings.Replace(groupIDPath+"/"+userInput.App.ArtifactId, "-", "", -1)
	return groupIDPath
}

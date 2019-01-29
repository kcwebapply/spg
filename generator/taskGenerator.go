package generator

import (
	"fmt"
	"go/build"
	"os"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

func GenerateTask(userInput parser.UserInput) {

	err := os.Mkdir(userInput.App.Name+path+"/task", 0777)
	if err != nil {
		fmt.Println("can'd make directory `task`.")
		os.Exit(0)
	}
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/src/main/java/task/task.java"
	writer := generateFile(userInput.App.Name + path + "/task/Task.java")
	content := getFormatFileContent(fileName)
	content = strings.Replace(content, "${schedule}", "\""+userInput.Task.Schedule+"\"", -1)
	content = strings.Replace(content, "${zone}", "\""+userInput.Task.Zone+"\"", -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

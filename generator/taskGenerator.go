package generator

import (
	"fmt"
	"os"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
	template "github.com/kcwebapply/spg/template"
)

// GenerateTaskClass generate @schedule annotated task class.
func GenerateTaskClass(userInput parser.UserInput) {
	schedule := userInput.Task.Schedule
	if &schedule == nil || schedule == "" {
		fmt.Println("please input schedule name.")
		return
	}

	taskPath := userInput.App.Name + path + "/" + getPathformatFromUserInput(userInput) + "/task"
	err := os.Mkdir(taskPath, 0777)
	if err != nil {
		fmt.Println("can'd make directory `task`.")
		os.Exit(0)
	}

	writer := generateFile(taskPath + "/Task.java")
	content := template.TASK
	content = strings.Replace(content, "${package}", getPackageformatFromUserInput(userInput), -1)
	content = strings.Replace(content, "${schedule}", "\""+userInput.Task.Schedule+"\"", -1)
	content = strings.Replace(content, "${zone}", "\""+userInput.Task.Zone+"\"", -1)
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

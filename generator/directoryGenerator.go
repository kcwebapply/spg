package generator

import (
	"fmt"
	"os"

	parser "github.com/kcwebapply/spg/parser"
)

const path = "/src/main/java"

// CreateDirectory initialize springBoot package.
func CreateDirectory(userInput parser.UserInput) {
	appName := userInput.App.Name
	appInfoPath := getPathformatFromUserInput(userInput)
	err := os.MkdirAll(appName+path+"/"+appInfoPath, 0777)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}

	err = os.MkdirAll(appName+"/src/test/java/"+appInfoPath, 0777)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}

	err = os.Mkdir(appName+"/src/main/resources", 0777)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(0)
	}
	os.Chmod(appName, 0777)
}

package generator

import (
	"fmt"
	"os"
)

// CreateDirectory initialize springBoot package.
func CreateDirectory(appName string) {
	err := os.MkdirAll(appName+path, 0777)
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

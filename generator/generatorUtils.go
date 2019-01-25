package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

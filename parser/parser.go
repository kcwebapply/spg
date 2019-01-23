package parser

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func Parse(fileName string) UserInput {
	var userInput UserInput
	_, err := toml.DecodeFile(fileName, &userInput)
	if err != nil {
		fmt.Println("err:", err)
	}
	return userInput
}

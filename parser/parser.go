package parser

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Parse method convert user's .toml file to UserInput struct.
func Parse(fileName string) UserInput {
	var userInput UserInput
	_, err := toml.DecodeFile(fileName, &userInput)
	if err != nil {
		fmt.Println("err:", err)
	}
	return userInput
}

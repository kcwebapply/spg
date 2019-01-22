package parser

import (
	"github.com/BurntSushi/toml"
)

func Parse(fileName string) UserInput {
	var userInput UserInput
	toml.DecodeFile(fileName, &userInput)
	return userInput
}

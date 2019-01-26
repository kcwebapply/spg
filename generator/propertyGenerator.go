package generator

import (
	"bufio"
	"fmt"

	parser "github.com/kcwebapply/spg/parser"
)

const datasourceDriverName = "spring.datasource.driver-class-name"
const datasourceURL = "spring.datasource.url"

// GeneratePropertiesFile =  touch application.properties
func GeneratePropertiesFile(userInput parser.UserInput) {
	writer := generateFile(userInput.App.Name + "/src/main/resources/application.properties")
	if userInput.Db.Driver != "" {
		setProperty(writer, datasourceURL, userInput.Db.Jdbc)
		setProperty(writer, datasourceDriverName, userInput.Db.Driver)
	}
	defer writer.Flush()
}

func setProperty(writer *bufio.Writer, key string, value string) {
	keyValue := fmt.Sprintf("%s=%s\n", key, value)
	writer.Write(([]byte)(keyValue))
}

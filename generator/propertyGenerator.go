package generator

import (
	"bufio"
	"fmt"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

const datasourceDriverName = "spring.datasource.driver-class-name"
const datasourceURL = "spring.datasource.url"

const mysqlDriver = "com.mysql.jdbc.Driver"
const postgreDriver = "org.postgresql.Driver"

// GeneratePropertiesFile touch application.properties
func GeneratePropertiesFile(userInput parser.UserInput) {
	writer := generateFile(userInput.App.Name + "/src/main/resources/application.properties")
	// db generation.
	if userInput.Db.Driver != "" || userInput.Db.Jdbc != "" {
		setProperty(writer, datasourceURL, userInput.Db.Jdbc)
		driver := userInput.Db.Driver
		if strings.Contains(driver, "postgres") {
			setProperty(writer, datasourceDriverName, postgreDriver)
		}

		if strings.Contains(driver, "mysql") {
			setProperty(writer, datasourceDriverName, mysqlDriver)
		}
	}
	defer writer.Flush()
}

func setProperty(writer *bufio.Writer, key string, value string) {
	keyValue := fmt.Sprintf("%s=%s\n", key, value)
	writer.Write(([]byte)(keyValue))
}

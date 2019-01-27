package generator

import (
	"fmt"
	"go/build"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

var jpaDependency = Dependency{GroupId: "org.springframework.boot", ArtifactId: "spring-boot-starter-data-jpa"}
var mySQLDependency = Dependency{GroupId: "mysql", ArtifactId: "mysql-connector-java"}
var postgresDependency = Dependency{GroupId: "org.postgresql", ArtifactId: "postgresql"}

const mysqlDriver = "com.mysql.jdbc.Driver"
const postgreDriver = "org.postgresql.Driver"

// GeneratePom touch pom.xml
func GeneratePom(userInput parser.UserInput) {

	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/pom.xml"
	content := getFormatFileContent(fileName)
	content = setProjectInfo(content, userInput)

	if userInput.Db.Driver != "" {
		driver := userInput.Db.Driver
		content = setDependency(content, jpaDependency)
		//org.mysql.Driver com.mysql.jdbc.Driver
		switch driver {
		case mysqlDriver:
			content = setDependency(content, mySQLDependency)
		case postgreDriver:
			content = setDependency(content, postgresDependency)
		}
	}
	writer := generateFile(userInput.App.Name + "/pom.xml")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

//set Project info
func setProjectInfo(content string, usetInput parser.UserInput) string {
	appName := usetInput.App.Name
	groupID := usetInput.App.GroupId
	artifactID := usetInput.App.ArtifactId
	contentText := strings.Replace(content, "${name}", appName, -1)
	contentText = strings.Replace(content, "${artifactId}", artifactID, -1)
	return strings.Replace(contentText, "${groupId}", groupID, -1)
}

// add dependency
func setDependency(content string, dependency Dependency) string {
	dependencyString := dependency.toString()
	return strings.Replace(content, "</dependencies>", dependencyString+"\n  </dependencies>", -1)
}

type Dependency struct {
	GroupId    string
	ArtifactId string
}

func (dependency *Dependency) toString() string {
	return fmt.Sprintf("  <dependency>\n        <groupId>%s</groupId>\n        <artifactId>%s</artifactId>\n    </dependency>", dependency.GroupId, dependency.ArtifactId)
}

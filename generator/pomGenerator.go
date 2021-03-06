package generator

import (
	"fmt"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
	template "github.com/kcwebapply/spg/template"
)

var jpaDependency = Dependency{GroupId: "org.springframework.boot", ArtifactId: "spring-boot-starter-data-jpa"}
var mySQLDependency = Dependency{GroupId: "mysql", ArtifactId: "mysql-connector-java"}
var postgresDependency = Dependency{GroupId: "org.postgresql", ArtifactId: "postgresql"}

// GeneratePom touch pom.xml
func GeneratePom(userInput parser.UserInput) {

	//fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/pom.xml"
	content := template.POM //getFormatFileContent(fileName)
	// set Project info on pom.xml
	content = setProjectInfo(content, userInput)
	// add Dependencies  on pom.xml
	content = setDependencies(content, userInput)

	writer := generateFile(userInput.App.Name + "/pom.xml")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

//set Project info
func setProjectInfo(content string, usetInput parser.UserInput) string {
	appName := usetInput.App.Name
	groupID := usetInput.App.GroupId
	artifactID := usetInput.App.ArtifactId
	springVersion := usetInput.App.SpringVersion
	javaVersion := usetInput.App.JavaVersion

	contentText := strings.Replace(content, "${name}", appName, -1)
	contentText = strings.Replace(contentText, "${artifactId}", artifactID, -1)
	contentText = strings.Replace(contentText, "${groupId}", groupID, -1)
	contentText = strings.Replace(contentText, "${springVersion}", springVersion, -1)
	contentText = strings.Replace(contentText, "${javaVersion}", javaVersion, -1)
	return contentText
}

func setDependencies(content string, userInput parser.UserInput) string {
	// dbSetting
	if userInput.Db.Driver != "" {
		driver := userInput.Db.Driver

		content = setDependency(content, jpaDependency)

		if strings.Contains(driver, "postgres") {
			content = setDependency(content, postgresDependency)
		}

		if strings.Contains(driver, "mysql") {
			content = setDependency(content, mySQLDependency)
		}
	}
	content = strings.Replace(content, "${dependencies}", "", -1)
	return content
}

// add dependency
func setDependency(content string, dependency Dependency) string {
	dependencyString := dependency.toString()
	return strings.Replace(content, "${dependencies}", dependencyString+"\n\n  ${dependencies}", -1)
}

type Dependency struct {
	GroupId    string
	ArtifactId string
}

func (dependency *Dependency) toString() string {
	return fmt.Sprintf("  <dependency>\n        <groupId>%s</groupId>\n        <artifactId>%s</artifactId>\n    </dependency>", dependency.GroupId, dependency.ArtifactId)
}

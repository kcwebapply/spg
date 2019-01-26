package generator

import (
	"fmt"
	"go/build"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

var jpaDependency = Dependency{GroupId: "org.springframework.boot", ArtifactId: "spring-boot-starter-data-jpa"}

// GeneratePom =  touch pom
func GeneratePom(userInput parser.UserInput) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/pom.xml"
	content := getFormatFileContent(fileName)
	if userInput.Db.Driver != "" {
		content = setDependency(content, jpaDependency)
	}
	writer := generateFile(userInput.App.Name + "/pom.xml")
	defer writer.Flush()
	writer.Write(([]byte)(content))
}

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

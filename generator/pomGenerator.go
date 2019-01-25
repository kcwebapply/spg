package generator

import (
	"fmt"
	"go/build"
	"strings"

	parser "github.com/kcwebapply/spg/parser"
)

var JpaDependency = Dependency{GroupId: "org.springframework.boot", ArtifactId: "spring-boot-starter-data-jpa"}

//"<dependency><groupId>org.springframework.boot</groupId><artifactId>spring-boot-starter-data-jpa</artifactId></dependency>"

// GeneratePom =  touch pom
func GeneratePom(userInput parser.UserInput) {
	fileName := build.Default.GOPATH + "/src/github.com/kcwebapply/spg/java/pom.xml"
	content := getFormatFileContent(fileName)
	if &userInput.Db != nil {
		content = setDependency(content, JpaDependency)
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
	Version    string
}

func (dependency *Dependency) toString() string {
	return fmt.Sprintf("  <dependency>\n        <groupId>%s</groupId>\n        <artifactId>%s</artifactId>\n        <version>%s</version>\n    </dependency>", dependency.GroupId, dependency.ArtifactId, dependency.Version)
}

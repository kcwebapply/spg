package parser

// UserInput is user's project setting.
type UserInput struct {
	App App `toml:App`
	Db  Db  `toml:Db`
}

// App is struct of user's application setting.
type App struct {
	Name          string `toml:name`
	GroupId       string `toml:groupId`
	ArtifactId    string `toml:artifactId`
	SpringVersion string `toml:springVersion`
	JavaVersion   string `toml:javaVersion`
}

//Db is struct of user's application setting.
type Db struct {
	Jdbc   string `toml:jdbc`
	Driver string `toml:driver`
	Table  string `toml:table`
}

type Batch struct {
}

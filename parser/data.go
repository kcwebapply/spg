package parser

type UserInput struct {
	App App `toml:App`
	Db  Db  `toml:Db`
}

type App struct {
	Name string `toml:name`
}

type Db struct {
	Jdbc  string `toml:jdbc`
	Table string `toml:table`
}

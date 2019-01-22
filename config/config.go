package config

import (
	"github.com/BurntSushi/toml"
)

func GetConfig() Config {
	var config Config
	toml.DecodeFile("config.toml", &config)
	return config
}

type Config struct {
	App App `toml:App`
}

type App struct {
	Name    string `toml:name`
	Version string `toml:version`
}

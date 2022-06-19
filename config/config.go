package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
	"golang.org/x/exp/slices"
)

const CURRENT_ENV = "test"

/*
	PROD yada TEST config ayarları için dev_config.json dosyası çoğaltılır.
	örnekler;
	dev_config.json
	test_config.json
	prod_config.json
*/

type Config struct {
	Env struct {
		PRODUCTION bool   `json:"production"`
		BASEURL    string `json:"baseurl"`
	}
	Credentials struct {
		USERNAME string `json:"username"`
		PASSWORD string `json:"password"`
	}
	Commands []Command
}

type Command struct {
	Name  string
	Value string
}

func GetConfig(params ...string) Config {
	conf := Config{}
	env := CURRENT_ENV
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	gonfig.GetConf(fileName, &conf)
	return conf
}

func GetCommand(commandName string) Command {
	if commandName == "" {
		panic("command name is empty")
	}
	commands := GetConfig().Commands
	idx := slices.IndexFunc(commands, func(command Command) bool { return command.Name == commandName })
	if idx == -1 {
		panic("command not found")
	}
	return commands[idx]
}

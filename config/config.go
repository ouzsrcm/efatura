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
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
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

func GetCommand(commandId string) Command {
	if commandId == "" {
		panic("command id cannot be null")
	}
	commands := GetConfig().Commands
	idx := slices.IndexFunc(commands, func(command Command) bool { return command.ID == commandId })
	if idx == -1 {
		panic("command not found")
	}
	return commands[idx]
}

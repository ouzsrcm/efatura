package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Config struct {
	Env struct {
		PRODUCTION bool
		PRODURL    string
		TESTURL    string
	} `yam:"ENV"`
	Credentials struct {
		USERNAME string
		PASSWORD string
	} `yaml:"credentials"`
}

func GetConfig(params ...string) Config {
	conf := Config{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	gonfig.GetConf(fileName, &conf)
	return conf
}

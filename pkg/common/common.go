package common

import (
	"invoicer/config"

	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func Configuration() config.Config {
	cfg := config.GetConfig()
	return cfg
}

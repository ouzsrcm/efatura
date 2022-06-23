package common

import (
	"invoicer/config"
	"time"

	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func Configuration() config.Config {
	cfg := config.GetConfig()
	return cfg
}

func ParseDate(date string) time.Time {
	layout := "2006-01-02"
	res, error := time.Parse(layout, date)
	if error != nil {
		panic(error)
	}
	return res
}

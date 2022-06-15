package main

import (
	"fmt"
	"invoicer/config"
)

func main() {
	cfg := config.GetConfig("dev")
	fmt.Println(cfg.Env.PRODURL)
}

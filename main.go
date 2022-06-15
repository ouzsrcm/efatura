package main

import (
	"fmt"
	"invoicer/config"
	"invoicer/pkg/invoice"
)

func main() {

	cfg := config.GetConfig("dev")
	fmt.Println(cfg.Env.PRODURL)

	token := invoice.GetToken("username", "password")
	fmt.Println(token)

}

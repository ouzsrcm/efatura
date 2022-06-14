package main

import (
	"fmt"
)

func main() {
	config := GetConfig("dev")
	fmt.Println(config.Env.PRODURL)
}

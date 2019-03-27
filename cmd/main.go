package main

import (
	"fmt"
	"idoc/config"
)

func main() {
	conf := config.Load("config.toml")
	fmt.Printf("%v", conf)
}

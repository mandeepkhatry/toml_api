package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func intializeRoutes(resourceLocation string) {
	if _, err := toml.DecodeFile(resourceLocation); err != nil {
		fmt.Println("Error")
	} else {

	}
}

func main() {
	resourceLocation := "resource.toml"
	intializeRoutes(resourceLocation)
}

package main

import (
	"./initializer"
)

func main() {
	resourceLocation := "resource.toml"
	initializer.IntializeRoutes(resourceLocation)
}

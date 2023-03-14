package main

import (
	"my2/config"
	"my2/unzip"
)

func main() {
	config.GetValue()
	unzip.ExtractZip(*(config.Config.Path))
}

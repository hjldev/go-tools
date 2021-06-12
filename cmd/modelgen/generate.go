package main

import (
	"flag"

	"github.com/hjldev/go-tools/internal/modelgen"
)

var ConfigPath = flag.String("config", "./config/modelgen.json", "comma-separated list of type names; must be set")

func main() {
	flag.Parse()
	modelgen.DbToGoStruct(*ConfigPath)
}

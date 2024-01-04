package main

//go:generate go run .

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	updateFlag = flag.Bool("update", false, "Update all specs.")

	clients = []string{"riot", "lol", "tft", "val", "lor"}

	SPECS_URLS = [][]string{
		{"http://www.mingweisamuel.com/riotapi-schema/openapi-3.0.0.json", "./specs/spec.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/queueTypes.json", "./specs/queueTypes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/gameTypes.json", "./specs/gameTypes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/gameModes.json", "./specs/gameModes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/routesTable.json", "./specs/routesTable.json"},
	}
)

func main() {
	flag.Parse()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths := strings.Split(path, "/")
	if paths[len(paths)-1] == "equinox" {
		err := os.Chdir("./codegen")
		if err != nil {
			panic(err)
		}
	}

	if *updateFlag || os.Getenv("UPDATE_SPECS") == "1" {
		fmt.Printf("Downloading specs...\n")
		for _, spec := range SPECS_URLS {
			err := DownloadAndSaveSpecs(spec)
			if err != nil {
				panic(err)
			}
		}
	}

	err = Compile()
	if err != nil {
		panic(err)
	}

	err = Format()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done\n")
}

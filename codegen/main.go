package main

//go:generate go run .

import (
	"flag"
	"fmt"
	"os"
)

var (
	updateFlag = flag.Bool("update", false, "Update all specs.")

	clients = []string{"riot", "lol", "tft", "val", "lor"}

	SPECS_URLS = [][]string{
		{"http://www.mingweisamuel.com/riotapi-schema/openapi-3.0.0.json", "./specs/spec.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/queues.json", "./specs/queues.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/queueTypes.json", "./specs/queueTypes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/gameTypes.json", "./specs/gameTypes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/gameModes.json", "./specs/gameModes.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/enums/maps.json", "./specs/maps.json"},
		{"http://www.mingweisamuel.com/riotapi-schema/routesTable.json", "./specs/routesTable.json"},
	}
)

func init() { flag.Parse() }

func main() {
	if *updateFlag || os.Getenv("UPDATE_SPECS") == "1" {
		fmt.Printf("Downloading specs...\n")
		for _, spec := range SPECS_URLS {
			err := DownloadAndSaveSpecs(spec)
			if err != nil {
				panic(err)
			}
		}
	}

	err := Compile()
	if err != nil {
		panic(err)
	}

	err = Format()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done\n")
}

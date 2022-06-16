//go:build integration
// +build integration

package integration

import (
	"fmt"
	"os"

	"github.com/Kyagara/equinox"
)

var (
	client *equinox.Equinox
)

func init() {
	key := os.Getenv("RIOT_GAMES_API_KEY")

	if key == "" {
		fmt.Println("RIOT_GAMES_API_KEY not found. Tests won't run.")
	} else {
		c, err := equinox.NewClient(key)

		if err != nil {
			fmt.Println(err)

			return
		}

		client = c
	}
}

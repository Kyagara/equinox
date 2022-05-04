# Equinox

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Kyagara/equinox)
[![Test Status](https://github.com/Kyagara/equinox/workflows/Tests/badge.svg)](https://github.com/Kyagara/equinox/actions?query=workflow%3Atests)

Equinox is a Riot Games API client written in golang with the goal of providing an easy to use interface to interact with all of the [Riot Games API](https://developer.riotgames.com/apis) endpoints.

#### Endpoints implemented:

-   [ ] Data Dragon
-   [x] Riot Account
-   [x] League of Legends
-   [x] Teamfight Tactics
-   [ ] Valorant
-   [x] Legends of Runeterra

## Example

```go
package main

import (
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/lol"
)

func main() {
	// For custom configurations, you can use NewClientWithConfig(),
	// you will need to provide an api.EquinoxConfig{} object.
	client, err := equinox.NewClient("RIOT_API_KEY")

	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}

	// Get this week's champion rotations.
	rotation, err := client.LOL.Champion.Rotations(lol.BR1)

	if err != nil {
		fmt.Println("error retrieving champion rotations: ", err)
		return
	}

	fmt.Printf("%+v\n", rotation)
	// &{FreeChampionIDs:[17 43 56 62 67 79 85 90 133 145 147 157 201 203 245 518]
	// FreeChampionIDsForNewPlayers:[222 254 427 82 131 147 54 17 18 37]
	// MaxNewPlayerLevel:10}
}
```

## Disclaimer

Equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

# Equinox

[![equinox release (latest SemVer)](https://img.shields.io/github/v/release/Kyagara/equinox?sort=semver)](https://github.com/Kyagara/equinox/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Kyagara/equinox)
[![Test Status](https://github.com/Kyagara/equinox/workflows/Tests/badge.svg)](https://github.com/Kyagara/equinox/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/Kyagara/equinox/branch/master/graph/badge.svg)](https://codecov.io/gh/Kyagara/equinox)

Equinox is a Riot Games API client written in golang with the goal of providing an easy to use interface to interact with all of the [Riot Games API](https://developer.riotgames.com/apis) endpoints and the Data Dragon service.

#### Implemented:

-   [ ] Data Dragon
-   [x] Riot Account
-   [x] League of Legends
-   [x] Teamfight Tactics
-   [x] Valorant
-   [x] Legends of Runeterra

## Installation

To install equinox you can either import it in a package:

```go
import "github.com/Kyagara/equinox"
```

and run `go get` without any parameters, or, if you want to use the latest version from this repo, you can use the following command:

```bash
go get github.com/Kyagara/equinox@master
```

## Usage

To access the diferent parts of the Riot Games API, create a new instance of the Equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
```
A client without a configuration struct comes with the default options:

```go
config := &api.EquinoxConfig{
	Key: "RIOT_API_KEY", // The API Key provided as a parameter.
	Cluster:  api.Americas, // Riot cluster, this should be one that is closest to you. Options available: Americas, Europe, Asia.
	LogLevel: api.FatalLevel, // The logging level, the FatalLevel effectively disables logging.
	Timeout:  10, // http.Client timeout in seconds.
	Retry:    true, // Retry if the API returns a 429 response.
}
```

A configuration struct can be used to create a custom Client using `equinox.NewClientWithConfig()`.

Now you can access different games endpoints by their abbreviations provided you have access to them. For example:

```go
// Can be accessed with a Development key.
summoner, err := client.LOL.Summoner.ByName(lol.BR1, "Loveable Senpai")

// May not be available in your policy.
recentMatches, err := client.VAL.Match.Recent(val.BR, val.CompetitiveQueue)
```

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

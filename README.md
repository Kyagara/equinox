<div align="center">
	<h1>equinox</h1>
	<p>
		<a href="https://github.com/Kyagara/equinox/releases">
			<img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=Version&style=flat-square"/>
		</a>  
		<a href="https://pkg.go.dev/github.com/Kyagara/equinox">
			<img src="https://img.shields.io/static/v1?label=Godoc&message=reference&color=blue&style=flat-square"/>
		</a>
		<a href="https://github.com/Kyagara/equinox/actions?query=workflow">
			<img src="https://img.shields.io/github/actions/workflow/status/Kyagara/equinox/ci.yaml?label=CI&style=flat-square"/>
		</a>
		<a href="https://codecov.io/gh/Kyagara/equinox">
			<img src="https://img.shields.io/codecov/c/github/Kyagara/equinox?style=flat-square"/>
		</a>
	</p>
	<p>
		<a href="#features">Features</a> •
		<a href="#todo">Todo</a> •
		<a href="#usage">Usage</a> •
		<a href="#example">Example</a> •
		<a href="#about">About</a> •
		<a href="#disclaimer">Disclaimer</a>
	</p>
</div>

## Features

- Riot APIs implemented:
  - Riot Account
  - League of Legends
  - Teamfight Tactics
  - Valorant
  - Legends of Runeterra
- Data Dragon and Community Dragon (Incomplete)
- Caching
- Retry on 429

## Todo

- More tests for the client
- Rate limit
- Improve DDragon/CDragon support
- Maybe more logging

## Usage

Get the library:

```bash
go get github.com/Kyagara/equinox
```

To access the diferent parts of the Riot Games API, create a new instance of the Equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
```

A client without a configuration comes with the default options:

```go
cacheConfig := bigcache.DefaultConfig(4 * time.Minute)
config := &api.EquinoxConfig{
	Key: "RIOT_API_KEY", // The API Key provided as a parameter.
	LogLevel: api.WARN_LOG_LEVEL, // The logging level, NOP_LOG_LEVEL disables logging.
	Timeout: 15, // http.Client timeout in seconds.
	Cache: cache.NewBigCache(cacheConfig), // The default client uses BigCache with an eviction time of 4 minutes.
	Retry: true, // Retries a request if the API returns a 429 response.
}
```

> A custom Client can be created using `equinox.NewClientWithConfig()`, requires an `&api.EquinoxConfig{}` struct.

> A different storage can be provided to the client using `cache.NewRedis()` or `cache.NewBigCache()`, passing nil in config.Cache disables caching.

Now you can access different games endpoints by their abbreviations. For example:

```go
// This method uses a lol.PlatformRoute. Can be accessed with a Development key.
summoner, err := client.LOL.SummonerV4.ByID(lol.BR1, "summoner_id")

// This method uses a api.RegionalRoute. Can be accessed with a Development key.
match, err := client.LOL.MatchV5.ByID(api.AMERICAS, "match_id")

// This method uses a api.RegionalRoute. Can be accessed with a Development key.
account, err := client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")

// This method uses a val.PlatformRoute. May not be available in your policy.
matches, err := client.VAL.MatchV1.Recent(val.BR, "competitive")
```

## Example

```go
package main

import (
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/clients/lol"
)

func main() {
	client, err := equinox.NewClient("RIOT_API_KEY")
	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}
	// Get this week's champion rotation.
	rotation, err := client.LOL.ChampionV3.Rotation(lol.BR1)
	if err != nil {
		fmt.Println("error retrieving champion rotation: ", err)
		return
	}
	fmt.Printf("%+v\n", rotation)
	// &{FreeChampionIDs:[17 43 56 62 67 79 85 90 133 145 147 157 201 203 245 518]
	// FreeChampionIDsForNewPlayers:[222 254 427 82 131 147 54 17 18 37]
	// MaxNewPlayerLevel:10}
}
```

## About

This is my first time developing and publishing an API client, I am constantly changing the project as I test and learn new things, please, check the commits for any breaking changes.

These projects helped me learn a lot:

- [go-github](https://github.com/google/go-github)
- [golio](https://github.com/KnutZuidema/golio)
- [Riven](https://github.com/MingweiSamuel/Riven)

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

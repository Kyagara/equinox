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

-   Riot APIs implemented:
    -   Riot Account
    -   League of Legends
    -   Teamfight Tactics
    -   Valorant
    -   Legends of Runeterra
-   Data Dragon
-   Caching
-   Rate Limiting

## Todo

-   New rate limiting implementation
-   Add Redis as an option for rate limit
-   Improve Data Dragon support

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
	Cluster: api.AmericasCluster, // Riot API cluster, use the cluster closest to you.
	LogLevel: api.FatalLevel, // The logging level, the FatalLevel provided effectively disables logging.
	Timeout: 15, // http.Client timeout in seconds.
	Cache: cache.NewBigCache(cacheConfig), // The default client uses BigCache with an eviction time of 4 minutes.
	Retry: true, // Retries a request if the API returns a 429 response.
	RateLimit: rate_limit.NewInternalRateLimit() // In-memory rate limit.
}
```

> A custom Client can be created using `equinox.NewClientWithConfig()`, requires an `&api.EquinoxConfig{}` struct.

> A different storage can be provided to the client using `cache.NewRedis()` or `cache.NewBigCache()`, passing nil in config.Cache disables caching.

Now you can access different games endpoints by their abbreviations. For example:

```go
// This method uses a lol.Region and a summoner name. Can be accessed with a Development key.
summoner, err := client.LOL.Summoner.ByName(lol.BR1, "Phanes")

// This method uses a lol.Route and a match ID. Can be accessed with a Development key.
summoner, err := client.LOL.Match.ByID(lol.Americas, "BR1_2530718601")

// The client.Cluster will be used as the region. Can be accessed with a Development key.
account, err := client.Riot.Account.ByPUUID("puuid")

// This method uses a val.Shard an a val.Queue. May not be available in your policy.
matches, err := client.VAL.Match.Recent(val.BR, val.CompetitiveQueue)
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
	// For custom configurations, you can use NewClientWithConfig(),
	// you will need to provide an &api.EquinoxConfig{}.
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

## About

This is my first time developing and publishing an API client, I am constantly changing the project as I test and learn new things, please, check the commits for any breaking changes.

These two projects helped me learn a lot:

-   [go-github](https://github.com/google/go-github)
-   [golio](https://github.com/KnutZuidema/golio)

## Disclaimer

Equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

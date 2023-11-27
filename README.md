<div align="center">
	<h1>equinox</h1>
	<img src="https://img.shields.io/github/go-mod/go-version/Kyagara/equinox?style=flat-square">
	<a href="https://github.com/Kyagara/equinox/tags"><img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=Version&style=flat-square"/></a>
	<a href="https://pkg.go.dev/github.com/Kyagara/equinox"><img src="https://img.shields.io/static/v1?label=Godoc&message=reference&color=blue&style=flat-square"/></a>
	<a href="https://codecov.io/gh/Kyagara/equinox"><img src="https://img.shields.io/codecov/c/github/Kyagara/equinox?style=flat-square"/></a>
	<p>
		<a href="#features">Features</a> •
		<a href="#todo">Todo</a> •
		<a href="#usage">Usage</a> •
		<a href="#example">Example</a>
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
- Rate limit (WIP)
- Caching
- Retry on 429

> equinox currently uses the proposed [jsonv2](https://github.com/go-json-experiment/json) package, read more about [here](https://github.com/golang/go/discussions/63397).

## Todo

- Rework retry, I believe the ratelimit is not respecting retry after header
- Properly use context
- Maybe create a custom BigCache config
- Fix issue with some ByAccessToken methods not being cached (dont want to use auth header as cache key)
- More tests for the rate limit
- More tests for the client
- Improve DDragon/CDragon support

## Usage

Get the library:

```bash
go get github.com/Kyagara/equinox
```

Create a new instance of the Equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
```

A client without a configuration comes with the default options:

- **Key**: The provided key.
- **LogLevel**: `WARN_LOG_LEVEL`. `NOP_LOG_LEVEL` disables logging.
- **HTTPClient**: `http.Client` with a timeout of 15 seconds.
- **Cache**: `BigCache` with an eviction time of 4 minutes.
- **Retry**: Retries a request 1 time if the API returns a 429 response.

```go
cacheConfig := bigcache.DefaultConfig(4 * time.Minute)
config := &api.EquinoxConfig{
	Key: "RIOT_API_KEY",
	LogLevel: api.WARN_LOG_LEVEL,
	HTTPClient: &http.Client{Timeout: 15 * time.Second},
	Cache: cache.NewBigCache(cacheConfig),
	Retry: 1,
}
```

> A custom Client can be created using `equinox.NewClientWithConfig()`, requires an `&api.EquinoxConfig{}` struct.

> A different storage can be provided to the client using `cache.NewRedis()` or `cache.NewBigCache()`, passing nil in config.Cache disables caching.

Readme more about the cache and rate limit implementations:

- [Cache](https://github.com/Kyagara/equinox/tree/master/cache)
- [Rate limit](https://github.com/Kyagara/equinox/tree/master/ratelimit)

Using different endpoints:

```go
// This method uses a lol.PlatformRoute. Can be accessed with a Development key.
summoner, err := client.LOL.SummonerV4.ByPUUID(lol.BR1, "summoner_puuid")

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

This is my first time developing and publishing an API client, I am constantly changing the project as I test and learn new things, please, check the commits for any breaking changes, this will change coming 1.0.0.

These projects helped me learn a lot:

- [go-github](https://github.com/google/go-github)
- [golio](https://github.com/KnutZuidema/golio)
- [Riven](https://github.com/MingweiSamuel/Riven)

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

<div align="center">
	<h1>equinox</h1>
	<img src="https://img.shields.io/github/go-mod/go-version/Kyagara/equinox?style=flat-square&label=go">
	<a href="https://github.com/Kyagara/equinox/tags"><img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=release&style=flat-square"/></a>
	<a href="https://pkg.go.dev/github.com/Kyagara/equinox"><img src="https://img.shields.io/static/v1?label=godoc&message=reference&color=blue&style=flat-square"/></a>
	<a href="https://codecov.io/gh/Kyagara/equinox"><img src="https://img.shields.io/codecov/c/github/Kyagara/equinox?style=flat-square&color=blue&label=coverage"/></a>
	<p>
		<a href="#usage">Usage</a> •
		<a href="#example">Example</a> •
		<a href="#todo">Todo</a> •
		<a href="#about">About</a>
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
- Retry `n` times

> equinox currently uses the proposed [jsonv2](https://github.com/go-json-experiment/json) package, read more about it [here](https://github.com/golang/go/discussions/63397).

## Usage

Get the library:

```bash
go get github.com/Kyagara/equinox # or: go get github.com/Kyagara/equinox@master
```

Create a new instance of the Equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
```

A client without a configuration comes with the default options:

- **Key**: The provided key.
- **LogLevel**: `zerolog.WarnLevel`. `zerolog.Disabled` disables logging.
- **Retries**: Retries a request n times if the API returns an error, defaults to 3.
- **HTTPClient**: `http.Client` with a timeout of 15 seconds.
- **Cache**: `BigCache` with an eviction time of 4 minutes.
- **RateLimit**: Internal rate limiter.

> A custom Client can be created using `equinox.NewClientWithConfig()`.

> A different storage can be provided to the client using `cache.NewRedis()` or `cache.NewBigCache()`, passing nil in config.Cache disables caching.

> You can disable rate limiting by passing `nil` in config.RateLimit, this can be useful if you have Equinox passing through a proxy that handles rate limiting.

> See [Cache](https://github.com/Kyagara/equinox/tree/master/cache) and [Rate limit](https://github.com/Kyagara/equinox/tree/master/ratelimit) for more details.

Using different endpoints:

```go
// This method uses a api.RegionalRoute. Can be accessed with a Development key.
match, err := client.LOL.MatchV5.ByID(api.AMERICAS, "match_id")

// This method uses a val.PlatformRoute. May not be available in your policy.
matches, err := client.VAL.MatchV1.Recent(val.BR, "competitive")

// Creating a request and executing it with ExecuteRaw which returns []byte but skips checking cache.
l := client.Internal.Logger("LOL_StatusV4_Platform")
req, err := client.Internal.Request(ctx, l, api.RIOT_API_BASE_URL_FORMAT, http.MethodGet, lol.BR1, "/lol/status/v4/platform-data", "", nil)
data, err :=client.Internal.ExecuteRaw(ctx, req)

// You can also interact with the cache.
data, err := client.Cache.Get("https://...")
err := client.Cache.Set("https://...", []byte("data"))
```

## Example

> For a slightly more advanced example, check out [lol-match-crawler](https://github.com/Kyagara/lol-match-crawler).

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

## Todo

- Maybe create a custom BigCache config
- More tests for the internal client and rate limit
- Improve DDragon/CDragon support

## About

This is my first time developing and publishing an API client, I am constantly changing the project as I test and learn new things, please, check the commits for any breaking changes, this will change coming 1.0.0.

These projects helped me learn a lot:

- [go-github](https://github.com/google/go-github)
- [golio](https://github.com/KnutZuidema/golio)
- [Riven](https://github.com/MingweiSamuel/Riven)

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

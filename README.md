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
- Data Dragon and Community Dragon - (Incomplete)
- Rate limit (Internal)
- Caching with [BigCache](https://github.com/allegro/bigcache) or [Redis](https://github.com/go-redis/redis)
- Logging with [zerolog](https://github.com/rs/zerolog)
- Exponential backoff

> equinox currently uses the proposed [jsonv2](https://github.com/go-json-experiment/json) package, read more about it [here](https://github.com/golang/go/discussions/63397).

## Usage

Get the library:

```bash
go get github.com/Kyagara/equinox # or: go get github.com/Kyagara/equinox@main
```

Create a new instance of the equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
// Or:
client, err := equinox.NewClientWithConfig(api.EquinoxConfig{})
```

> Cache and Rate Limit can be disabled by passing `nil`, disabling rate limit can be useful if you have equinox passing through a proxy that handles rate limiting. See [Cache](https://github.com/Kyagara/equinox/tree/main/cache) and [Rate limit](https://github.com/Kyagara/equinox/tree/main/ratelimit) for more details.

A default equinox client comes with the default options:

- **Key**: The provided Riot API key.
- **HTTPClient**: `http.Client` with a timeout of 15 seconds.
- **Cache**: `BigCache` with an eviction time of 4 minutes.
- **RateLimit**: Internal rate limiter with a limit usage factor of 1.0 and interval overhead of 1 second.
- **Logger**: api.Logger object with `zerolog.WarnLevel`. Will log if rate limited or when retrying a request (before waiting).
- **Retry**: api.Retry object with a limit of 3 and jitter of 500 milliseconds.

Using different endpoints:

```go
// Contexts without a deadline will block if rate limited, waiting for the rate limited buckets to reset.
// If a deadline is set, there will be checks done before any block to see if waiting would exceed that deadline.
ctx := context.Background()

// This method uses an api.RegionalRoute. Can be accessed with a Development key.
match, err := client.LOL.MatchV5.ByID(ctx, api.AMERICAS, "BR1_2...")

// This method uses an val.PlatformRoute. May not be available in your policy.
matches, err := client.VAL.MatchV1.Recent(ctx, val.BR, "competitive")

// Interacting with DDragon.
version, err := client.DDragon.Version.Latest(ctx)
champion, err := client.CDragon.Champion.ByName(ctx, version, "Aatrox")

// Interacting with the cache.
data, err := client.Cache.Get("https://...")
err := client.Cache.Set("https://...", data)

// Using ExecuteRaw which returns []byte but skips checking cache.
l := client.Internal.Logger("LOL_StatusV4_Platform")
urlComponents := []string{"https://", lol.BR1, api.RIOT_API_BASE_URL_FORMAT, "/lol/status/v4/platform-data"}
req, err := client.Internal.Request(ctx, l, http.MethodGet, urlComponents, "", nil)
data, err := client.Internal.ExecuteRaw(ctx, req)
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
	ctx := context.Background()
	rotation, err := client.LOL.ChampionV3.Rotation(ctx, lol.BR1)
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

- Add Redis store for rate limit, using a lua script
- Maybe the context usage throughout the project could be improved
- Maybe add more options to customize the rate limiter
- More tests for the internal client and rate limit
- Improve DDragon/CDragon support

## About

This is my first time developing and publishing an API client, any feedback is appreciated.

equinox started as a way for me to learn go, many projects helped me to understand the inner workings of an API client and go itself, here are some of them, check them out.

- [go-github](https://github.com/google/go-github)
- [golio](https://github.com/KnutZuidema/golio)

Projects not written in go:

- [Riven](https://github.com/MingweiSamuel/Riven)
- [riot-api](https://github.com/fightmegg/riot-api) and [riot-rate-limiter](https://github.com/fightmegg/riot-rate-limiter)
- [galeforce](https://github.com/bcho04/galeforce)

A rewrite of Riven's [code generation](https://github.com/MingweiSamuel/Riven/tree/v/2.x.x/riven/srcgen) is used in [equinox](https://github.com/Kyagara/equinox/tree/main/codegen).

If cloning the project, make sure to have a `go.work` file in the root of the project, if you don't, you will get an error when trying to run `go generate`.

```
go 1.21

use (
	.
	./codegen
)
```

If you have any questions, feel free to [open an issue](https://github.com/Kyagara/equinox/issues) or contact me.

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

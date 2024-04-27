<div align="center">
	<h1>equinox</h1>
	<img src="https://img.shields.io/github/go-mod/go-version/Kyagara/equinox?style=flat-square&label=go">
	<a href="https://github.com/Kyagara/equinox/tags"><img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=release&style=flat-square"/></a>
	<a href="https://pkg.go.dev/github.com/Kyagara/equinox"><img src="https://img.shields.io/static/v1?label=godoc&message=reference&color=blue&style=flat-square"/></a>
	<a href="https://codecov.io/gh/Kyagara/equinox"><img src="https://img.shields.io/codecov/c/github/Kyagara/equinox?style=flat-square&color=blue&label=coverage"/></a>
	<p>
		<a href="https://github.com/Kyagara/equinox/wiki">Wiki</a> •
		<a href="#example">Example</a> •
		<a href="#todo">Todo</a>
	</p>
</div>

## Features

- Riot APIs implemented:
  - Riot Account
  - League of Legends
  - Teamfight Tactics
  - Valorant
  - Legends of Runeterra
- Rate limit (Internal)
- Caching with [BigCache](https://github.com/allegro/bigcache) or [Redis](https://github.com/go-redis/redis)
- Logging with [zerolog](https://github.com/rs/zerolog)
- Exponential backoff

> [!NOTE]
> equinox currently uses the proposed [jsonv2](https://github.com/go-json-experiment/json), read more about it [here](https://github.com/golang/go/discussions/63397).

Check the [Wiki](https://github.com/Kyagara/equinox/wiki) for more information about the library.

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

- Maybe the context usage throughout the project could be improved
- Improve key concatenation in Redis methods
- New RateLimit interface
  - Similar to the cache interface
  - Add Redis store, maybe using a lua script
  - Maybe add more options (presets?) to customize the rate limiter
- More tests for the internal client and rate limit

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

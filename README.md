<div align="center">
	<h1>equinox</h1>
	<img src="https://img.shields.io/github/go-mod/go-version/Kyagara/equinox?style=flat-square&label=go">
	<a href="https://github.com/Kyagara/equinox/tags"><img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=release&style=flat-square"/></a>
	<a href="https://pkg.go.dev/github.com/Kyagara/equinox/v2"><img src="https://img.shields.io/static/v1?label=godoc&message=reference&color=blue&style=flat-square"/></a>
	<a href="https://codecov.io/gh/Kyagara/equinox"><img src="https://img.shields.io/codecov/c/github/Kyagara/equinox?style=flat-square&color=blue&label=coverage"/></a>
	<p>
		<a href="https://github.com/Kyagara/equinox/wiki">Wiki</a> •
		<a href="#example">Example</a> •
		<a href="#todo">Todo</a> •
		<a href="#versioning">Versioning</a>
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

	"github.com/Kyagara/equinox/v2/"
	"github.com/Kyagara/equinox/v2/clients/lol"
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
- Maybe add and move `Logger` to the `Cache` and `RateLimit` interfaces
- Maybe endpoint methods `int32`/`int64` parameters should be changed to just `int`
- Add checks for duration of tests that include any WaitN/any blocking
- Add more integration tests
- RateLimit
  - Add Redis store, using a lua script
  - Maybe add more options (presets?) to customize the rate limiter
  - Try to reduce amount of method arguments

## Versioning

Breaking changes in the library itself (removal/rename of methods from `InternalClient`/`Cache`/`RateLimit`) will require a major version (n.x.x) bump, fixes will occur in a patch (x.x.n).

The Riot API does not follow semver and changes often, breaking changes such as removal of endpoints methods or even entire endpoints will require a **minor** version (x.n.x) bump.

In `go`, new major versions of a library are an annoyance, requiring adding/changing **every** import path to the new version, so I want to keep them to a minimum. For example, `github.com/Kyagara/equinox/internal/client` will be `github.com/Kyagara/equinox/v2/internal/client`, then `github.com/Kyagara/equinox/v3/internal/client`...

Older versions are not supported (versions before v1 are outright broken and shouldn't be used), always keep the library up-to-date with the latest version.

## Disclaimer

equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

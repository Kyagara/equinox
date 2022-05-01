# Equinox

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Kyagara/equinox)
[![Test Status](https://github.com/Kyagara/equinox/workflows/Tests/badge.svg)](https://github.com/Kyagara/equinox/actions?query=workflow%3Atests)

This library is NOT production ready, expect breaking changes.

Only some League of Legends endpoints are currently implemented.

This project is a first for me since I don't have much experience in publishing a library, this is pretty much a practice tool for me to improve my knowledge in CI/CD using Github Actions, golang and tests.

I am avoiding using other packages like [resty](https://github.com/go-resty/resty) instead of the `net/http` package go provides to improve my golang knowledge, currently using [testify](https://github.com/stretchr/testify) and [gock](https://github.com/h2non/gock) for tests.

## Example

```go
package main

import (
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
)

func main() {
	// For custom configurations, you can use NewClientWithConfig(),
	// you will need to provide an api.EquinoxConfig{} object.
	client, err := equinox.NewClient("RIOT_API_KEY")

	if err != nil {
		fmt.Println("error creating client", err)
		return
	}

	// Get champion rotations.
	rotation, err := client.LOL.Champion.Rotations(api.LOLRegionBR1)

	if err != nil {
		fmt.Println("error retrieving champion rotations", err)
		return
	}

	fmt.Printf("%+v\n", rotation)
}
```

## TODO

#### DTOs

DTOs are found inside their respective endpoint implementation, however, in some cases where an endpoint has multiple methods, the files quickly become a mess as the DTOs occupy a large portion of the file, maybe they could be implemented in another module, however I am not sure on how to organize them since I plan to support other endpoints from other Riot games.

#### Improve tests

First time doings mocks so I decided to use [gock](https://github.com/h2non/gock) just to get things out of the ground, not sure if this implementation of tests with mocks is good.

Tests could benefit from using Parallel().

#### Improve Logging

I don't believe the current logging and 'debugging mode' is done right or that its 'done' at all to be honest, there might be places where I should be logging something but I am not.

## Disclaimer

Equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.

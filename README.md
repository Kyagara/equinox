# Equinox

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/Kyagara/equinox)
[![Test Status](https://github.com/Kyagara/equinox/workflows/Tests/badge.svg)](https://github.com/Kyagara/equinox/actions?query=workflow%3Atests)

This shouldn't be used in any production environment, this is just a practice tool for me to learn CI/CD using Github Actions and tests in golang.

I was recommended [Alex Pliutau](https://www.youtube.com/watch?v=evorkFq3Y5k)'s video on youtube and got curious about learning other things, I decided to make a client for the Riot Games API since I am more familiar with it.

I am avoiding using other packages like [resty](https://github.com/go-resty/resty) instead of the `net/http` package go provides to improve my golang knowledge, currently just using [testify](https://github.com/stretchr/testify) for tests.

## Example

```go
package main

import (
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
)

func main() {
	// Or you can use NewClientWithDebug() to enable debugging,
	// this will print requests before and after they are sent
	client, err := equinox.NewClient("RIOT_API_KEY")

	if err != nil {
		fmt.Println("error creating client,", err)
		return
	}

	// Get Free Champion rotation
	champions, err := client.LOL.Champion.FreeRotation(api.LOLRegionBR1)

	if err != nil {
		fmt.Println("error retrieving champions,", err)
		return
	}

	fmt.Printf("%+v\n", champions)
}
```

## TODO

### Tests

I am not sure if tests are 'good enough', I am just checking if errors are `Nil` and the response is `NotNil` using `testify`.

### Logging

I don't believe the current logging and 'debugging mode' is done right or that its 'done' at all to be honest, there might be places where I should be logging something but I am not.

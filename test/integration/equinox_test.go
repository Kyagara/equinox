//go:build integration
// +build integration

package integration

import (
	"os"
	"time"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/rs/zerolog"
)

var (
	client *equinox.Equinox
)

func init() {
	key := os.Getenv("RIOT_GAMES_API_KEY")
	if key == "" {
		panic("RIOT_GAMES_API_KEY not set")
	}

	// Default client with pretty logging and lower MaxRetries
	config := equinox.DefaultConfig(key)
	config.Retry = api.Retry{MaxRetries: 1, Jitter: 500 * time.Millisecond}
	config.Logger = api.Logger{Pretty: true, Level: zerolog.TraceLevel, EnableTimestamp: true, EnableConfigurationLogging: true}

	cache, err := equinox.DefaultCache()
	if err != nil {
		panic(err)
	}

	ratelimit := equinox.DefaultRateLimit()

	client, err = equinox.NewCustomClient(config, nil, cache, ratelimit)
	if err != nil {
		panic(err)
	}
}

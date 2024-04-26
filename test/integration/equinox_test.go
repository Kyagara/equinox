//go:build integration
// +build integration

package integration

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/allegro/bigcache/v3"
)

var (
	client *equinox.Equinox
)

func init() {
	key := os.Getenv("RIOT_GAMES_API_KEY")
	if key == "" {
		panic("RIOT_GAMES_API_KEY not set")
	}

	// Default client with a test logger and lower MaxRetries

	ctx := context.Background()
	cacheConfig := bigcache.DefaultConfig(4 * time.Minute)
	cacheConfig.Verbose = false
	cache, err := cache.NewBigCache(ctx, cacheConfig)
	if err != nil {
		panic(err)
	}

	config := api.EquinoxConfig{
		Key: key,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Cache:     cache,
		RateLimit: ratelimit.NewInternalRateLimit(0.99, time.Second),
		Retry:     api.Retry{MaxRetries: 1, Jitter: 500 * time.Millisecond},
		Logger:    util.TestLogger(),
	}

	client, err = equinox.NewClientWithConfig(config)
	if err != nil {
		panic(err)
	}
}

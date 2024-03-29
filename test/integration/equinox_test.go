//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
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

var onlyDataDragon = false

func checkIfOnlyDataDragon(t *testing.T) {
	if onlyDataDragon {
		t.Skip()
	}
}

func init() {
	key := os.Getenv("RIOT_GAMES_API_KEY")
	if key == "" || key == "RGAPI..." || key == "RGAPI-TEST" {
		fmt.Println("RIOT_GAMES_API_KEY not found. Only Data Dragon tests will run.")
		onlyDataDragon = true
		key = "RGAPI-TEST"
	}

	ctx := context.Background()
	cache, err := cache.NewBigCache(ctx, bigcache.DefaultConfig(4*time.Minute))
	if err != nil {
		fmt.Println(err)
		return
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

	client = equinox.NewClientWithConfig(config)
}

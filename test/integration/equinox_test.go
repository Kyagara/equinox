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
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/require"
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
	config := &api.EquinoxConfig{
		Key:      key,
		LogLevel: api.DEBUG_LOG_LEVEL,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Retry: 1,
		Cache: cache,
	}
	c, err := equinox.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client = c
}

func TestClientCache(t *testing.T) {
	checkIfOnlyDataDragon(t)
	rotations, err := client.LOL.ChampionV3.Rotation(lol.BR1)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, rotations, "expecting non-nil rotations")
	start := time.Now()
	cachedRotations, err := client.LOL.ChampionV3.Rotation(lol.BR1)
	duration := int(time.Since(start).Seconds())
	require.Equal(t, rotations, cachedRotations)
	require.Nil(t, err, "expecting nil error")
	if duration >= 2 {
		err = fmt.Errorf("request took more than 1s, took %ds, request not cached, check logs to see if this is an error", duration)
		require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))
	}
}

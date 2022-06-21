//go:build integration
// +build integration

package integration

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/stretchr/testify/require"
)

var (
	client *equinox.Equinox
)

func init() {
	key := os.Getenv("RIOT_GAMES_API_KEY")

	if key == "" {
		fmt.Println("RIOT_GAMES_API_KEY not found. Tests won't run.")
	} else {
		config := &api.EquinoxConfig{
			Key:       key,
			Cluster:   api.AmericasCluster,
			LogLevel:  api.DebugLevel,
			Timeout:   10,
			TTL:       240,
			Retry:     true,
			RateLimit: true,
		}

		c, err := equinox.NewClientWithConfig(config)

		if err != nil {
			fmt.Println(err)

			return
		}

		client = c
	}
}

func TestClientCache(t *testing.T) {
	rotations, err := client.LOL.Champion.Rotations(lol.BR1)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, rotations, "expecting non-nil rotations")

	start := time.Now()
	cachedRotations, err := client.LOL.Champion.Rotations(lol.BR1)
	duration := int(time.Since(start).Seconds())

	require.Equal(t, rotations, cachedRotations)

	require.Nil(t, err, "expecting nil error")

	if duration >= 2 {
		err = fmt.Errorf("request took more than 1s, took %ds, request not cached, check logs to see if this is an error", duration)

		require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))
	}
}

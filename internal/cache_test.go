package internal_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestNewCache(t *testing.T) {
	cache := internal.NewCache()

	require.NotNil(t, cache, "expecting non-nil Cache")
}

func TestCacheHit(t *testing.T) {
	internalClient := internal.NewInternalClient(&api.EquinoxConfig{
		Key:      "RIOT_API_KEY",
		Cluster:  api.AmericasCluster,
		LogLevel: api.DebugLevel,
		Timeout:  10,
		TTL:      120,
		Retry:    true,
	})

	client := lol.NewLOLClient(internalClient)

	j := []byte("{\"freeChampionIds\":[2,12,27,29,32,35,42,54,72,78,84,98,136,164,223,777],\"freeChampionIdsForNewPlayers\":[222,254,427,82,131,147,54,17,18,37],\"maxNewPlayerLevel\":10}")

	data := &lol.ChampionRotationsDTO{}

	err := json.Unmarshal(j, data)

	require.Nil(t, err, "expecting nil error")

	defer gock.Off()

	gock.New(fmt.Sprintf(api.BaseURLFormat, "br1")).
		Get(lol.ChampionURL).
		Reply(200).JSON(data)

	gotData, gotErr := client.Champion.Rotations(lol.BR1)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	require.Equal(t, data, gotData, fmt.Sprintf("want data %v, got %v", data, gotData))

	gotCache, gotErr := client.Champion.Rotations(lol.BR1)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	require.Equal(t, data, gotCache, fmt.Sprintf("want data %v, got %v", data, gotCache))
}

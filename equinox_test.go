package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/allegro/bigcache/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestNewEquinoxClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		key     string
	}{
		{
			name: "success",
			want: &equinox.Equinox{},
			key:  "RGAPI-KEY",
		},
		{
			name:    "nil key",
			wantErr: fmt.Errorf("API Key not provided"),
			key:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClient(test.key)

			if test.name != "success" {
				require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

				if test.wantErr == nil {
					require.Equal(t, test.want, client)
				}
			} else {
				require.NotNil(t, client, "expecting non-nil Client")
			}
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		config  *api.EquinoxConfig
	}{
		{
			name:   "success",
			want:   &equinox.Equinox{},
			config: internal.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: fmt.Errorf("equinox configuration not provided"),
			config:  nil,
		},
		{
			name:    "api key nil",
			wantErr: fmt.Errorf("API Key not provided"),
			config: &api.EquinoxConfig{
				LogLevel: api.DebugLevel,
				Timeout:  15,
				Retry:    true,
			},
		},
		{
			name:    "cluster nil",
			wantErr: fmt.Errorf("cluster not provided"),
			config: &api.EquinoxConfig{
				Key:      "RGAPI-KEY",
				LogLevel: api.DebugLevel,
				Timeout:  15,
				Retry:    true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)

			if test.name != "success" {
				require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

				if test.wantErr == nil {
					require.Equal(t, test.want, client)
				}
			} else {
				require.NotNil(t, client, "expecting non-nil Client")
			}
		})
	}
}

func TestEquinoxClientClearCache(t *testing.T) {
	cache, err := cache.NewBigCache(bigcache.DefaultConfig(4 * time.Minute))

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	config := internal.NewTestEquinoxConfig()

	config.Cache = cache

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(t, err, "expecting nil error")

	account := &riot.AccountDTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	delay := 2 * time.Second

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
		Get(fmt.Sprintf(riot.AccountByPUUIDURL, "puuid")).
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AmericasCluster)).
		Get(fmt.Sprintf(riot.AccountByPUUIDURL, "puuid")).
		Reply(403).
		JSON(account).Delay(delay)

	gotData, gotErr := client.Riot.Account.ByPUUID("puuid")

	require.Equal(t, account, gotData)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start := time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration := int(time.Since(start).Seconds())

	require.Equal(t, account, gotData, fmt.Sprintf("want data %v, got %v", account, gotData))

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		t.Error(fmt.Errorf("request took more than 1s, took %ds, request not cached", duration))
	}

	client.ClearCache()

	start = time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)

	assert.Equal(t, api.ForbiddenError, gotErr, fmt.Sprintf("want err %v, got %v", api.ForbiddenError, gotErr))
}

// Never done a benchmark in go
// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkCachedSummonerByName-16 74191 16400 ns/op 10135 B/op 54 allocs/op
// BenchmarkCachedSummonerByName-16 73783 16356 ns/op 10160 B/op 54 allocs/op
// BenchmarkCachedSummonerByName-16 71601 16173 ns/op 10300 B/op 54 allocs/op
// BenchmarkCachedSummonerByName-16 70710 17575 ns/op 10359 B/op 54 allocs/op
// BenchmarkCachedSummonerByName-16 68031 17968 ns/op 10547 B/op 54 allocs/op
func BenchmarkCachedSummonerByName(b *testing.B) {
	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(fmt.Sprintf(lol.SummonerByNameURL, "Loveable Senpai")).
		Persist().
		Reply(200).
		JSON(&lol.SummonerDTO{
			ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
			AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
			PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
			Name:          "Loveable Senpai",
			ProfileIconID: 1386,
			RevisionDate:  1657211888000,
			SummonerLevel: 68,
		})

	client, err := equinox.NewClient("RGAPI-")

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.Summoner.ByName(lol.BR1, "Loveable Senpai")

		require.Nil(b, err)

		require.Equal(b, "Loveable Senpai", data.Name)
	}
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkSummonerByName-16 33925 31289 ns/op 11525 B/op 112 allocs/op
// BenchmarkSummonerByName-16 38095 31338 ns/op 11650 B/op 113 allocs/op
// BenchmarkSummonerByName-16 37455 31614 ns/op 11907 B/op 114 allocs/op
// BenchmarkSummonerByName-16 37741 31735 ns/op 11907 B/op 114 allocs/op
// BenchmarkSummonerByName-16 36934 32410 ns/op 12420 B/op 115 allocs/op
func BenchmarkSummonerByName(b *testing.B) {
	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(fmt.Sprintf(lol.SummonerByNameURL, "Loveable Senpai")).
		Persist().
		Reply(200).
		JSON(&lol.SummonerDTO{
			ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
			AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
			PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
			Name:          "Loveable Senpai",
			ProfileIconID: 1386,
			RevisionDate:  1657211888000,
			SummonerLevel: 68,
		})

	config := internal.NewTestEquinoxConfig()

	config.LogLevel = api.FatalLevel

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.Summoner.ByName(lol.BR1, "Loveable Senpai")

		require.Nil(b, err)

		require.Equal(b, "Loveable Senpai", data.Name)
	}
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkCachedDataDragonRealm-16 64609 18678 ns/op 10809 B/op 61 allocs/op
// BenchmarkCachedDataDragonRealm-16 64401 18629 ns/op 10826 B/op 61 allocs/op
// BenchmarkCachedDataDragonRealm-16 63915 18615 ns/op 10866 B/op 61 allocs/op
// BenchmarkCachedDataDragonRealm-16 61941 18680 ns/op 11033 B/op 61 allocs/op
// BenchmarkCachedDataDragonRealm-16 64098 19215 ns/op 10850 B/op 61 allocs/op
func BenchmarkCachedDataDragonRealm(b *testing.B) {
	realm := &data_dragon.RealmData{
		N: struct {
			Item        string "json:\"item\""
			Rune        string "json:\"rune\""
			Mastery     string "json:\"mastery\""
			Summoner    string "json:\"summoner\""
			Champion    string "json:\"champion\""
			ProfileIcon string "json:\"profileicon\""
			Map         string "json:\"map\""
			Language    string "json:\"language\""
			Sticker     string "json:\"sticker\""
		}{
			Item:        "12.13.1",
			Rune:        "7.23.1",
			Mastery:     "7.23.1",
			Summoner:    "12.13.1",
			Champion:    "12.13.1",
			ProfileIcon: "12.13.1",
			Map:         "12.13.1",
			Language:    "12.13.1",
			Sticker:     "12.13.1",
		},
		V:              "12.13.1",
		L:              "en_US",
		Cdn:            "https://ddragon.leagueoflegends.com/cdn",
		Dd:             "12.13.1",
		Lg:             "12.13.1",
		CSS:            "12.13.1",
		Profileiconmax: 28,
		Store:          nil,
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	client, err := equinox.NewClient("RGAPI-")

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.DataDragon.Realm.ByName(data_dragon.NA)

		require.Nil(b, err)

		require.Equal(b, "en_US", data.L)
	}
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkDataDragonRealm-16 37357 31601 ns/op 11468 B/op 118 allocs/op
// BenchmarkDataDragonRealm-16 37443 32366 ns/op 11594 B/op 119 allocs/op
// BenchmarkDataDragonRealm-16 36427 32602 ns/op 11851 B/op 120 allocs/op
// BenchmarkDataDragonRealm-16 36283 32618 ns/op 11851 B/op 120 allocs/op
// BenchmarkDataDragonRealm-16 35736 33285 ns/op 12365 B/op 121 allocs/op
func BenchmarkDataDragonRealm(b *testing.B) {
	realm := &data_dragon.RealmData{
		N: struct {
			Item        string "json:\"item\""
			Rune        string "json:\"rune\""
			Mastery     string "json:\"mastery\""
			Summoner    string "json:\"summoner\""
			Champion    string "json:\"champion\""
			ProfileIcon string "json:\"profileicon\""
			Map         string "json:\"map\""
			Language    string "json:\"language\""
			Sticker     string "json:\"sticker\""
		}{
			Item:        "12.13.1",
			Rune:        "7.23.1",
			Mastery:     "7.23.1",
			Summoner:    "12.13.1",
			Champion:    "12.13.1",
			ProfileIcon: "12.13.1",
			Map:         "12.13.1",
			Language:    "12.13.1",
			Sticker:     "12.13.1",
		},
		V:              "12.13.1",
		L:              "en_US",
		Cdn:            "https://ddragon.leagueoflegends.com/cdn",
		Dd:             "12.13.1",
		Lg:             "12.13.1",
		CSS:            "12.13.1",
		Profileiconmax: 28,
		Store:          nil,
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	config := internal.NewTestEquinoxConfig()

	config.LogLevel = api.FatalLevel

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.DataDragon.Realm.ByName(data_dragon.NA)

		require.Nil(b, err)

		require.Equal(b, "en_US", data.L)
	}
}

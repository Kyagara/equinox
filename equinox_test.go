package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/rate_limit"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			key:  "RGAPI-TEST",
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
				Key:      "RGAPI-TEST",
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
	config, err := equinox.DefaultConfig("RGAPI-TEST")

	config.LogLevel = api.DebugLevel
	config.Retry = false

	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

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
		Reply(404).
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

	gotErr = client.Cache.Clear()

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start = time.Now()
	gotData, gotErr = client.Riot.Account.ByPUUID("puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)

	assert.Equal(t, api.ErrNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrNotFound, gotErr))
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkCachedSummonerByName-16 106034 11073 ns/op 5766 B/op 36 allocs/op
// BenchmarkCachedSummonerByName-16  99432 11565 ns/op 5977 B/op 36 allocs/op
// BenchmarkCachedSummonerByName-16 104392 11258 ns/op 5816 B/op 36 allocs/op
// BenchmarkCachedSummonerByName-16 106242 11265 ns/op 5760 B/op 36 allocs/op
// BenchmarkCachedSummonerByName-16 103720 11193 ns/op 5837 B/op 36 allocs/op
func BenchmarkCachedSummonerByName(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerDTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(fmt.Sprintf(lol.SummonerByNameURL, "Phanes")).
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.Summoner.ByName(lol.BR1, "Phanes")

		require.Nil(b, err)

		require.Equal(b, "Phanes", data.Name)
	}
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkSummonerByName-16 59298 20242 ns/op 5240 B/op 76 allocs/op
// BenchmarkSummonerByName-16 58087 20730 ns/op 5367 B/op 77 allocs/op
// BenchmarkSummonerByName-16 56473 21340 ns/op 5622 B/op 78 allocs/op
// BenchmarkSummonerByName-16 52909 21888 ns/op 5622 B/op 78 allocs/op
// BenchmarkSummonerByName-16 54087 22303 ns/op 6135 B/op 79 allocs/op
func BenchmarkSummonerByName(b *testing.B) {
	b.ReportAllocs()

	summoner := &lol.SummonerDTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	headers := map[string]string{}

	headers[rate_limit.AppRateLimitHeader] = "1300:2,1300:60"
	headers[rate_limit.AppRateLimitCountHeader] = "1:2,5:60"

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(fmt.Sprintf(lol.SummonerByNameURL, "Phanes")).
		Persist().
		Reply(200).
		JSON(summoner).SetHeaders(headers)

	config := internal.NewTestEquinoxConfig()

	config.LogLevel = api.NopLevel
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.LOL.Summoner.ByName(lol.BR1, "Phanes")

		require.Nil(b, err)

		require.Equal(b, "Phanes", data.Name)
	}
}

// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkCachedDataDragonRealm-16 84229 14301 ns/op 6692 B/op 44 allocs/op
// BenchmarkCachedDataDragonRealm-16 83107 14246 ns/op 6746 B/op 44 allocs/op
// BenchmarkCachedDataDragonRealm-16 83066 14168 ns/op 6748 B/op 44 allocs/op
// BenchmarkCachedDataDragonRealm-16 83488 14247 ns/op 6727 B/op 44 allocs/op
// BenchmarkCachedDataDragonRealm-16 83391 14222 ns/op 6732 B/op 44 allocs/op
func BenchmarkCachedDataDragonRealm(b *testing.B) {
	b.ReportAllocs()

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
		CDN:            "https://ddragon.leagueoflegends.com/cdn",
		DD:             "12.13.1",
		LG:             "12.13.1",
		CSS:            "12.13.1",
		ProfileIconMax: 28,
		Store:          nil,
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	client, err := equinox.NewClient("RGAPI-TEST")

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
// BenchmarkDataDragonRealm-16 50066 23434 ns/op 5298 B/op 82 allocs/op
// BenchmarkDataDragonRealm-16 48817 23760 ns/op 5425 B/op 83 allocs/op
// BenchmarkDataDragonRealm-16 47798 24149 ns/op 5678 B/op 84 allocs/op
// BenchmarkDataDragonRealm-16 48996 25465 ns/op 5678 B/op 84 allocs/op
// BenchmarkDataDragonRealm-16 45928 25454 ns/op 6191 B/op 85 allocs/op
func BenchmarkDataDragonRealm(b *testing.B) {
	b.ReportAllocs()

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
		CDN:            "https://ddragon.leagueoflegends.com/cdn",
		DD:             "12.13.1",
		LG:             "12.13.1",
		CSS:            "12.13.1",
		ProfileIconMax: 28,
		Store:          nil,
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	config := internal.NewTestEquinoxConfig()

	config.LogLevel = api.NopLevel
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.DataDragon.Realm.ByName(data_dragon.NA)

		require.Nil(b, err)

		require.Equal(b, "en_US", data.L)
	}
}

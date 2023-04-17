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

// Never done a benchmark in go
// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 7 2700 Eight-Core Processor
// BenchmarkCachedSummonerByName-16 75765 15572 ns/op 9640 B/op 48 allocs/op
// BenchmarkCachedSummonerByName-16 76281 15520 ns/op 9611 B/op 48 allocs/op
// BenchmarkCachedSummonerByName-16 77018 15946 ns/op 9569 B/op 48 allocs/op
// BenchmarkCachedSummonerByName-16 73472 16236 ns/op 9780 B/op 48 allocs/op
// BenchmarkCachedSummonerByName-16 68229 17395 ns/op 10132 B/op 48 allocs/op
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
// BenchmarkSummonerByName-16 48676 24291 ns/op 7663 B/op 82 allocs/op
// BenchmarkSummonerByName-16 48651 24363 ns/op 7789 B/op 83 allocs/op
// BenchmarkSummonerByName-16 45241 24926 ns/op 8047 B/op 84 allocs/op
// BenchmarkSummonerByName-16 47240 25171 ns/op 8046 B/op 84 allocs/op
// BenchmarkSummonerByName-16 45874 25856 ns/op 8560 B/op 85 allocs/op
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

	config.LogLevel = api.FatalLevel
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
// BenchmarkCachedDataDragonRealm-16 69696 17329 ns/op 10028 B/op 55 allocs/op
// BenchmarkCachedDataDragonRealm-16 63756 17504 ns/op 10479 B/op 55 allocs/op
// BenchmarkCachedDataDragonRealm-16 68504 17458 ns/op 10113 B/op 55 allocs/op
// BenchmarkCachedDataDragonRealm-16 68016 18454 ns/op 10148 B/op 55 allocs/op
// BenchmarkCachedDataDragonRealm-16 60730 19914 ns/op 10743 B/op 55 allocs/op
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
// BenchmarkDataDragonRealm-16 44347 28088 ns/op 7664 B/op 89 allocs/op
// BenchmarkDataDragonRealm-16 45283 27041 ns/op 7791 B/op 90 allocs/op
// BenchmarkDataDragonRealm-16 38968 27811 ns/op 8049 B/op 91 allocs/op
// BenchmarkDataDragonRealm-16 40266 27107 ns/op 8049 B/op 91 allocs/op
// BenchmarkDataDragonRealm-16 42764 27618 ns/op 8561 B/op 92 allocs/op
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

	config.LogLevel = api.FatalLevel
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)

	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		data, err := client.DataDragon.Realm.ByName(data_dragon.NA)

		require.Nil(b, err)

		require.Equal(b, "en_US", data.L)
	}
}

package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
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
			name: "nil key",
			key:  "",
			want: &equinox.Equinox{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClient(test.key)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))
			require.NotNil(t, client, "expecting non-nil Client")
			require.NotNil(t, client.Cache, "expecting non-nil Client")
			require.NotNil(t, client.LOL, "expecting nil Client")
			require.NotNil(t, client.LOR, "expecting nil Client")
			require.NotNil(t, client.TFT, "expecting nil Client")
			require.NotNil(t, client.VAL, "expecting nil Client")
			require.NotNil(t, client.Riot, "expecting nil Client")
			require.NotNil(t, client.CDragon, "expecting non-nil Client")
			require.NotNil(t, client.DDragon, "expecting non-nil Client")
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
			name:    "no cache",
			want:    &equinox.Equinox{},
			config:  internal.NewTestEquinoxConfig(),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

			if test.wantErr == nil {
				require.NotNil(t, client, "expecting non-nil client")
			}

			if test.name == "no cache" {
				require.Equal(t, client.Cache.TTL, time.Duration(0), "expecting cache disabled")
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
	require.Equal(t, client.Cache.TTL, 4*time.Minute, "expecting cache enabled")

	account := &riot.AccountV1DTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	delay := 2 * time.Second

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.BaseURLFormat, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(404).
		JSON(account).Delay(delay)

	gotData, gotErr := client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	require.Equal(t, account, gotData)
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start := time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration := int(time.Since(start).Seconds())

	require.Equal(t, account, gotData, fmt.Sprintf("want data %v, got %v", account, gotData))
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		t.Error(fmt.Errorf("request took more than 1s, took %ds, request not cached", duration))
	}

	gotErr = client.Cache.Clear()
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start = time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)
	require.Equal(t, api.ErrNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrNotFound, gotErr))
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkCachedSummonerByName-16  113149 10367 ns/op 5543 B/op 34 allocs/op
BenchmarkCachedSummonerByName-16  106800 10656 ns/op 5720 B/op 34 allocs/op
BenchmarkCachedSummonerByName-16  109754 10439 ns/op 5635 B/op 34 allocs/op
BenchmarkCachedSummonerByName-16  110697 10672 ns/op 5609 B/op 34 allocs/op
BenchmarkCachedSummonerByName-16  100819 11005 ns/op 5907 B/op 34 allocs/op
*/
func BenchmarkCachedSummonerByName(b *testing.B) {
	b.ReportAllocs()
	summoner := &lol.SummonerV4DTO{
		Id:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-name/Phanes").
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.Nil(b, err)
	for i := 0; i < b.N; i++ {
		data, err := client.LOL.SummonerV4.ByName(lol.BR1, "Phanes")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkSummonerByName-16 67662 17969 ns/op 5169 B/op 71 allocs/op
BenchmarkSummonerByName-16 65629 17823 ns/op 5293 B/op 72 allocs/op
BenchmarkSummonerByName-16 64528 18342 ns/op 5549 B/op 73 allocs/op
BenchmarkSummonerByName-16 64230 18831 ns/op 5550 B/op 73 allocs/op
BenchmarkSummonerByName-16 61237 19108 ns/op 6062 B/op 74 allocs/op
*/
func BenchmarkSummonerByName(b *testing.B) {
	b.ReportAllocs()
	summoner := &lol.SummonerV4DTO{
		Id:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-name/Phanes").
		Persist().
		Reply(200).
		JSON(summoner)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.NopLevel
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)
	for i := 0; i < b.N; i++ {
		data, err := client.LOL.SummonerV4.ByName(lol.BR1, "Phanes")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkCachedDataDragonRealm-16  81088 14627 ns/op 6823 B/op 42 allocs/op
BenchmarkCachedDataDragonRealm-16  82737 13190 ns/op 6740 B/op 42 allocs/op
BenchmarkCachedDataDragonRealm-16  83905 13898 ns/op 6684 B/op 42 allocs/op
BenchmarkCachedDataDragonRealm-16  82736 14826 ns/op 6740 B/op 42 allocs/op
BenchmarkCachedDataDragonRealm-16  83233 14297 ns/op 6716 B/op 42 allocs/op
*/
func BenchmarkCachedDataDragonRealm(b *testing.B) {
	b.ReportAllocs()
	realm := &ddragon.RealmData{
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
		Get(fmt.Sprintf(ddragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.Nil(b, err)
	for i := 0; i < b.N; i++ {
		data, err := client.DDragon.Realm.ByName(ddragon.NA)
		require.Nil(b, err)
		require.Equal(b, "en_US", data.L)
	}
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkDataDragonRealm-16 56840 21223 ns/op 5255 B/op 79 allocs/op
BenchmarkDataDragonRealm-16 54490 21758 ns/op 5382 B/op 80 allocs/op
BenchmarkDataDragonRealm-16 54898 21924 ns/op 5638 B/op 81 allocs/op
BenchmarkDataDragonRealm-16 54536 22074 ns/op 5638 B/op 81 allocs/op
BenchmarkDataDragonRealm-16 53311 23353 ns/op 6155 B/op 82 allocs/op
*/
func BenchmarkDataDragonRealm(b *testing.B) {
	b.ReportAllocs()
	realm := &ddragon.RealmData{
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
		Get(fmt.Sprintf(ddragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.NopLevel
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)
	for i := 0; i < b.N; i++ {
		data, err := client.DDragon.Realm.ByName(ddragon.NA)
		require.Nil(b, err)
		require.Equal(b, "en_US", data.L)
	}
}

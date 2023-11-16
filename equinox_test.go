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
	emptyKeyConfig := internal.NewTestEquinoxConfig()
	emptyKeyConfig.Key = ""
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
	config.LogLevel = api.NOP_LOG_LEVEL
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

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
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

func TestPutClient(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := lol.NewLOLClient(internalClient)
	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Put("/lol/tournament/v5/codes/tournamentCode").
		Reply(200)

	err = client.TournamentV5.UpdateCode(api.AMERICAS, nil, "tournamentCode")
	require.Nil(t, err, "expecting nil error")
}

/*
goos: windows
goarch: amd64
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkCachedSummonerByName-16 119008 10058 ns/op 4976 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 118884  9912 ns/op 4979 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 118033  9800 ns/op 4999 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 116996  9861 ns/op 5025 B/op 33 allocs/op
BenchmarkCachedSummonerByName-16 111648  9993 ns/op 5163 B/op 33 allocs/op
*/
func BenchmarkCachedSummonerByName(b *testing.B) {
	b.ReportAllocs()
	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
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
BenchmarkSummonerByName-16 61848 19319 ns/op 4727 B/op 69 allocs/op
BenchmarkSummonerByName-16 61212 20062 ns/op 4858 B/op 70 allocs/op
BenchmarkSummonerByName-16 58365 20639 ns/op 5111 B/op 71 allocs/op
BenchmarkSummonerByName-16 55460 20802 ns/op 5112 B/op 71 allocs/op
BenchmarkSummonerByName-16 55996 21634 ns/op 5625 B/op 72 allocs/op
*/
func BenchmarkSummonerByName(b *testing.B) {
	b.ReportAllocs()
	summoner := &lol.SummonerV4DTO{
		ID:            "5kIdR5x9LO0pVU_v01FtNVlb-dOws-D04GZCbNOmxCrB7A",
		AccountID:     "NkJ3FK5BQcrpKtF6Rj4PrAe9Nqodd2rwa5qJL8kJIPN_BkM",
		PUUID:         "6WQtgEvp61ZJ6f48qDZVQea1RYL9akRy7lsYOIHH8QDPnXr4E02E-JRwtNVE6n6GoGSU1wdXdCs5EQ",
		Name:          "Phanes",
		ProfileIconID: 1386,
		RevisionDate:  1657211888000,
		SummonerLevel: 68,
	}

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
		Get("/lol/summoner/v4/summoners/by-name/Phanes").
		Persist().
		Reply(200).
		JSON(summoner)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
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
BenchmarkCachedDataDragonRealm-16 94222 12113 ns/op 5513 B/op 39 allocs/op
BenchmarkCachedDataDragonRealm-16 99193 11967 ns/op 5334 B/op 39 allocs/op
BenchmarkCachedDataDragonRealm-16 96424 12064 ns/op 5431 B/op 39 allocs/op
BenchmarkCachedDataDragonRealm-16 99111 12129 ns/op 5337 B/op 39 allocs/op
BenchmarkCachedDataDragonRealm-16 83335 12517 ns/op 5981 B/op 39 allocs/op
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

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
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
BenchmarkDataDragonRealm-16 55206 21476 ns/op 4104 B/op 73 allocs/op
BenchmarkDataDragonRealm-16 54673 22166 ns/op 4230 B/op 74 allocs/op
BenchmarkDataDragonRealm-16 52680 22500 ns/op 4487 B/op 75 allocs/op
BenchmarkDataDragonRealm-16 52269 22853 ns/op 4487 B/op 75 allocs/op
BenchmarkDataDragonRealm-16 52078 23503 ns/op 5002 B/op 76 allocs/op
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

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
		Get(fmt.Sprintf(ddragon.RealmURL, "na")).
		Persist().
		Reply(200).
		JSON(realm)

	config := internal.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)
	for i := 0; i < b.N; i++ {
		data, err := client.DDragon.Realm.ByName(ddragon.NA)
		require.Nil(b, err)
		require.Equal(b, "en_US", data.L)
	}
}

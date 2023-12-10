package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/h2non/gock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// This revealed problems with multiple limits in a bucket, only the first bucket was being respected.
func BenchmarkParallelRateLimit(b *testing.B) {
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

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1, "")).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).SetHeaders(map[string]string{
		ratelimit.APP_RATE_LIMIT_HEADER:       "20:1,50:3,80:5",
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER: "1:1,1:3,1:5",
	}).
		JSON(summoner)

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelCachedSummonerByPUUID-16 267324 4487 ns/op 2539 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 269991 4540 ns/op 2526 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 269116 4480 ns/op 2531 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 246655 4589 ns/op 2644 B/op 17 allocs/op
BenchmarkParallelCachedSummonerByPUUID-16 207062 7285 ns/op 2905 B/op 17 allocs/op
*/
func BenchmarkParallelCachedSummonerByPUUID(b *testing.B) {
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

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1, "")).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelSummonerByPUUID-16 151017  7424 ns/op 2805 B/op 36 allocs/op
BenchmarkParallelSummonerByPUUID-16 130140  8964 ns/op 2934 B/op 37 allocs/op
BenchmarkParallelSummonerByPUUID-16 106862 10016 ns/op 3192 B/op 38 allocs/op
BenchmarkParallelSummonerByPUUID-16 106148 10652 ns/op 3194 B/op 38 allocs/op
BenchmarkParallelSummonerByPUUID-16  91861 11961 ns/op 3711 B/op 39 allocs/op
*/
func BenchmarkParallelSummonerByPUUID(b *testing.B) {
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

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1, "")).
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
			require.NoError(b, err)
			require.Equal(b, "Phanes", data.Name)
		}
	})
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelDDragonRealms-16 121444 10430 ns/op 4006 B/op 45 allocs/op
BenchmarkParallelDDragonRealms-16 110125 11687 ns/op 4295 B/op 46 allocs/op
BenchmarkParallelDDragonRealms-16  97659 12349 ns/op 4322 B/op 46 allocs/op
BenchmarkParallelDDragonRealms-16  92355 13589 ns/op 4874 B/op 47 allocs/op
BenchmarkParallelDDragonRealms-16  82082 13934 ns/op 4907 B/op 47 allocs/op
*/
func BenchmarkParallelDDragonRealms(b *testing.B) {
	b.ReportAllocs()

	realms := &ddragon.RealmData{
		N: ddragon.RealmN{
			Item:        "13.24.1",
			Rune:        "7.23.1",
			Mastery:     "7.23.1",
			Summoner:    "13.24.1",
			Champion:    "13.24.1",
			ProfileIcon: "13.24.1",
			Map:         "13.24.1",
			Language:    "13.24.1",
			Sticker:     "13.24.1",
		},
		V:              "13.24.1",
		L:              "en_US",
		CDN:            "https://ddragon.leagueoflegends.com/cdn",
		DD:             "13.24.1",
		LG:             "13.24.1",
		CSS:            "13.24.1",
		ProfileIconMax: 28,
		Store:          nil,
	}

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "/realms/na.json", "")).
		Get("").
		Persist().
		Reply(200).
		JSON(realms)

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.DDragon.Realm.ByName(ctx, ddragon.NA)
			require.NoError(b, err)
			require.Equal(b, "13.24.1", data.V)
		}
	})
}

/*
goos: linux
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkParallelMatchListByPUUID-16 43246 26421 ns/op 12137 B/op 131 allocs/op
BenchmarkParallelMatchListByPUUID-16 42696 27607 ns/op 12335 B/op 132 allocs/op
BenchmarkParallelMatchListByPUUID-16 37814 28060 ns/op 12595 B/op 133 allocs/op
BenchmarkParallelMatchListByPUUID-16 41376 28777 ns/op 12687 B/op 133 allocs/op
BenchmarkParallelMatchListByPUUID-16 40714 29963 ns/op 13244 B/op 134 allocs/op
*/
func BenchmarkParallelMatchListByPUUID(b *testing.B) {
	b.ReportAllocs()

	list := []string{
		"BR1_2744215970",
		"BR1_2744215971",
		"BR1_2744215972",
	}

	gock.New("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/puuid/ids?count=10&queue=420&type=RANKED_SOLO_5x5").
		Get("").
		Persist().
		Reply(200).
		JSON(list)

	config := util.NewTestEquinoxConfig()
	config.LogLevel = zerolog.WarnLevel
	config.Retries = 3

	client, err := equinox.NewClientWithConfig(config)
	require.NoError(b, err)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx := context.Background()
			data, err := client.LOL.MatchV5.ListByPUUID(ctx, api.AMERICAS, "puuid", -1, -1, 420, string(lol.RANKED_SOLO_5X5), -1, 10)
			require.NoError(b, err)
			require.Equal(b, "BR1_2744215970", data[0])
		}
	})
}

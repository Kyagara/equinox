package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/test/util"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkCachedSummonerByPUUID-16 121608 8800 ns/op 4211 B/op 20 allocs/op
BenchmarkCachedSummonerByPUUID-16 122340 8553 ns/op 4195 B/op 20 allocs/op
BenchmarkCachedSummonerByPUUID-16 140514 8479 ns/op 3839 B/op 20 allocs/op
BenchmarkCachedSummonerByPUUID-16 137011 8501 ns/op 3900 B/op 20 allocs/op
BenchmarkCachedSummonerByPUUID-16 138278 8637 ns/op 3878 B/op 20 allocs/op
*/
// This version and the non cached version are used to estimate how the cache impacts performance.
func BenchmarkCachedSummonerByPUUID(b *testing.B) {
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
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	client, err := equinox.NewClient("RGAPI-TEST")
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkSummonerByPUUID-16 103070 11914 ns/op 2987 B/op 41 allocs/op
BenchmarkSummonerByPUUID-16  97792 12154 ns/op 3117 B/op 42 allocs/op
BenchmarkSummonerByPUUID-16  92349 12345 ns/op 3374 B/op 43 allocs/op
BenchmarkSummonerByPUUID-16  90201 12424 ns/op 3375 B/op 43 allocs/op
BenchmarkSummonerByPUUID-16  89172 13025 ns/op 3887 B/op 44 allocs/op
*/
func BenchmarkSummonerByPUUID(b *testing.B) {
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
		Get("/lol/summoner/v4/summoners/by-puuid/puuid").
		Persist().
		Reply(200).
		JSON(summoner)

	config := util.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = true

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		data, err := client.LOL.SummonerV4.ByPUUID(ctx, lol.BR1, "puuid")
		require.Nil(b, err)
		require.Equal(b, "Phanes", data.Name)
	}
}

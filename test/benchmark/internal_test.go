package benchmark

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/jarcoal/httpmock"
)

/*
Simulating an endpoint method

goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkInternals-16 301467 3775 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 287043 3790 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 303676 3782 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 289886 3777 ns/op 1402 B/op 17 allocs/op
BenchmarkInternals-16 290647 3792 ns/op 1402 B/op 17 allocs/op
*/
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client, err := internal.NewInternalClient(config)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger := client.Logger("LOL_SummonerV4_ByPUUID")
		ctx := context.Background()
		urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		if err != nil {
			b.Fatal(err)
		}
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		if err != nil {
			b.Fatal(err)
		}
		if data != "response" {
			b.Fatalf("data != response, got: %s", data)
		}
	}
}

/*
equinoxReq.Request.URL.String() causes 3 allocs, the alloc results are not modified to account for that

goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkRequest-16 724022 1602 ns/op 648 B/op 7 allocs/op
BenchmarkRequest-16 751404 1597 ns/op 648 B/op 7 allocs/op
BenchmarkRequest-16 631826 1594 ns/op 648 B/op 7 allocs/op
BenchmarkRequest-16 693741 1616 ns/op 648 B/op 7 allocs/op
BenchmarkRequest-16 633535 1611 ns/op 648 B/op 7 allocs/op
*/
func BenchmarkRequest(b *testing.B) {
	b.ReportAllocs()

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client, err := internal.NewInternalClient(config)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()
	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		if err != nil {
			b.Fatal(err)
		}
		if equinoxReq.Request.URL.String() != "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid" {
			b.Fatalf("URL != https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid, got: %s", equinoxReq.Request.URL.String())
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecute-16 427688 2455 ns/op 873 B/op 13 allocs/op
BenchmarkExecute-16 451112 2455 ns/op 873 B/op 13 allocs/op
BenchmarkExecute-16 468093 2439 ns/op 873 B/op 13 allocs/op
BenchmarkExecute-16 440035 2450 ns/op 873 B/op 13 allocs/op
BenchmarkExecute-16 460554 2440 ns/op 873 B/op 13 allocs/op
*/
func BenchmarkExecute(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client, err := internal.NewInternalClient(config)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data string
		err = client.Execute(ctx, equinoxReq, &data)
		if err != nil {
			b.Fatal(err)
		}
		if data != "response" {
			b.Fatalf("data != response, got: %s", data)
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkExecuteBytes-16 518600 2072 ns/op 1368 B/op 13 allocs/op
BenchmarkExecuteBytes-16 582710 2033 ns/op 1368 B/op 13 allocs/op
BenchmarkExecuteBytes-16 519804 2037 ns/op 1368 B/op 13 allocs/op
BenchmarkExecuteBytes-16 525337 2151 ns/op 1368 B/op 13 allocs/op
BenchmarkExecuteBytes-16 558279 2050 ns/op 1368 B/op 13 allocs/op
*/
func BenchmarkExecuteBytes(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()
	client, err := internal.NewInternalClient(config)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "br1", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/", "puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := client.ExecuteBytes(ctx, equinoxReq)
		if err != nil {
			b.Fatal(err)
		}
		if len(data) != 10 {
			b.Fatalf("data length != 10, got: %d", len(data))
		}
	}
}

/*
goos: linux - WSL2
goarch: amd64
pkg: github.com/Kyagara/equinox/test/benchmark
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkURLWithAuthorizationHash-16 2008862 597.8 ns/op 216 B/op 5 allocs/op
BenchmarkURLWithAuthorizationHash-16 2007656 592.6 ns/op 216 B/op 5 allocs/op
BenchmarkURLWithAuthorizationHash-16 2016522 594.7 ns/op 216 B/op 5 allocs/op
BenchmarkURLWithAuthorizationHash-16 1999180 593.0 ns/op 216 B/op 5 allocs/op
BenchmarkURLWithAuthorizationHash-16 1991668 603.5 ns/op 216 B/op 5 allocs/op
*/
func BenchmarkURLWithAuthorizationHash(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://br1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := util.NewTestEquinoxConfig()
	config.Logger = equinox.DefaultLogger()
	config.Retry = equinox.DefaultRetry()

	req := &http.Request{
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
		Header: http.Header{},
	}

	equinoxReq := api.EquinoxRequest{Request: req}
	equinoxReq.URL = req.URL.String()
	req.Header.Set("Authorization", "7267ee00-5696-47b8-9cae-8db3d49c8c33")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		url := internal.GetURLWithAuthorizationHash(equinoxReq)
		if url != "http://example.com/path-45da11db1ebd17ee0c32aca62e08923ea4f15590058ff1e15661bc13ed33df9d" {
			b.Fatalf("URL != http://example.com/path-45da11db1ebd17ee0c32aca62e08923ea4f15590058ff1e15661bc13ed33df9d, got: %s", url)
		}
	}
}

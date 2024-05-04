package benchmark

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/jarcoal/httpmock"
)

// Simulating an endpoint method.
func BenchmarkInternals(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger := client.Logger("LOL_SummonerV4_ByPUUID")
		ctx := context.Background()
		urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
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

func BenchmarkInternalRequest(b *testing.B) {
	b.ReportAllocs()

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()
	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
		if err != nil {
			b.Fatal(err)
		}
		if equinoxReq.URL != "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid" {
			b.Fatalf("URL != https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid, got: %s", equinoxReq.Request.URL.String())
		}
	}
}

func BenchmarkInternalExecute(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/puuid"}
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

func BenchmarkInternalExecuteBytes(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

	config := equinox.DefaultConfig("RGAPI-TEST")
	client, err := internal.NewInternalClient(config, nil, nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	logger := client.Logger("LOL_SummonerV4_ByPUUID")
	ctx := context.Background()

	urlComponents := []string{"https://", "kr", api.RIOT_API_BASE_URL_FORMAT, "/lol/summoner/v4/summoners/by-puuid/", "puuid"}
	equinoxReq, err := client.Request(ctx, logger, http.MethodGet, urlComponents, "summoner-v4.getByPUUID", nil)
	if err != nil {
		b.Fatal(err)
	}

	res := []byte(`"response"`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := client.ExecuteBytes(ctx, equinoxReq)
		if err != nil {
			b.Fatal(err)
		}
		if !bytes.Equal(data, res) {
			b.Fatal("data != response")
		}
	}
}

func BenchmarkInternalURLWithAuthorizationHash(b *testing.B) {
	b.ReportAllocs()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/puuid",
		httpmock.NewStringResponder(200, `"response"`))

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

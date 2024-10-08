package benchmark

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/jarcoal/httpmock"
)

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

	var data string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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

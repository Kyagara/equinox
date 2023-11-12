package ddragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestChampionAllChampions(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := ddragon.NewDDragonClient(internalClient)
	json := &ddragon.DDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    map[string]ddragon.ChampionData{},
	}
	data := map[string]ddragon.ChampionData{}

	tests := []struct {
		name    string
		code    int
		want    *ddragon.DDragonMetadata
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: json,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
				Get(fmt.Sprintf(ddragon.ChampionsURL, "1.0", ddragon.PtBR)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.AllChampions("1.0", ddragon.PtBR)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				require.Equal(t, data, gotData)
			}
		})
	}
}

func TestChampionByName(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := ddragon.NewDDragonClient(internalClient)

	json := &ddragon.DDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    map[string]ddragon.FullChampionData{},
	}

	json.Data.(map[string]ddragon.FullChampionData)["JarvanIV"] = ddragon.FullChampionData{}

	data := &ddragon.FullChampionData{}

	tests := []struct {
		name    string
		code    int
		want    *ddragon.DDragonMetadata
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: json,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
				Get(fmt.Sprintf(ddragon.ChampionURL, "1.0", ddragon.PtBR, "JarvanIV")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.ByName("1.0", ddragon.PtBR, "JarvanIV")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, data, gotData)
			}
		})
	}
}

func TestChampionAllChampionsBadlyFormattedJson(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := ddragon.NewDDragonClient(internalClient)

	json := &ddragon.DDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    "{bad:json}",
	}

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
		Get(fmt.Sprintf(ddragon.ChampionsURL, "1.0", ddragon.PtBR)).
		Reply(200).
		JSON(json)

	gotData, gotErr := client.Champion.AllChampions("1.0", ddragon.PtBR)

	require.Nil(t, gotData, "expecting nil data")

	error_string := "cannot unmarshal string into Go value of type map[string]ddragon.ChampionData"

	require.ErrorContains(t, gotErr, error_string)
}

func TestChampionByNameBadlyFormattedJson(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := ddragon.NewDDragonClient(internalClient)

	json := &ddragon.DDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    "{bad:json}",
	}

	gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
		Get(fmt.Sprintf(ddragon.ChampionURL, "1.0", ddragon.PtBR, "JarvanIV")).
		Reply(200).
		JSON(json)

	gotData, gotErr := client.Champion.ByName("1.0", ddragon.PtBR, "JarvanIV")

	require.Nil(t, gotData, "expecting nil data")

	error_string := "cannot unmarshal string into Go value of type map[string]ddragon.FullChampionData"

	require.ErrorContains(t, gotErr, error_string)
}

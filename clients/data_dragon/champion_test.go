package data_dragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChampionAllChampions(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	json := &data_dragon.DataDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    map[string]*data_dragon.ChampionData{},
	}

	data := map[string]*data_dragon.ChampionData{}

	tests := []struct {
		name    string
		code    int
		want    *data_dragon.DataDragonMetadata
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
			gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
				Get(fmt.Sprintf(data_dragon.ChampionsURL, "1.0", data_dragon.PtBR)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.AllChampions("1.0", data_dragon.PtBR)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, data, gotData)
			}
		})
	}
}

func TestChampionByName(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	json := &data_dragon.DataDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    map[string]*data_dragon.ChampionData{},
	}

	json.Data.(map[string]*data_dragon.ChampionData)["JarvanIV"] = &data_dragon.ChampionData{}

	data := &data_dragon.ChampionData{}

	tests := []struct {
		name    string
		code    int
		want    *data_dragon.DataDragonMetadata
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
			gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
				Get(fmt.Sprintf(data_dragon.ChampionURL, "1.0", data_dragon.PtBR, "JarvanIV")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Champion.ByName("1.0", data_dragon.PtBR, "JarvanIV")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, data, gotData)
			}
		})
	}
}

func TestChampionAllChampionsBadlyFormattedJson(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	json := &data_dragon.DataDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    "{bad:json}",
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.ChampionsURL, "1.0", data_dragon.PtBR)).
		Reply(200).
		JSON(json)

	gotData, gotErr := client.Champion.AllChampions("1.0", data_dragon.PtBR)

	require.Nil(t, gotData, "expecting nil data")

	error_string := "cannot unmarshal string into Go value of type map[string]*data_dragon.ChampionData"

	require.ErrorContains(t, gotErr, error_string)
}

func TestChampionByNameBadlyFormattedJson(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	json := &data_dragon.DataDragonMetadata{
		Type:    "",
		Format:  "",
		Version: "",
		Data:    "{bad:json}",
	}

	gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
		Get(fmt.Sprintf(data_dragon.ChampionURL, "1.0", data_dragon.PtBR, "JarvanIV")).
		Reply(200).
		JSON(json)

	gotData, gotErr := client.Champion.ByName("1.0", data_dragon.PtBR, "JarvanIV")

	require.Nil(t, gotData, "expecting nil data")

	error_string := "cannot unmarshal string into Go value of type map[string]*data_dragon.ChampionData"

	require.ErrorContains(t, gotErr, error_string)
}

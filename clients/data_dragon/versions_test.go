package data_dragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestVersionLatest(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]string{"1.0", "0.9"},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
				Get(data_dragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Version.Latest()

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				ver := "1.0"
				assert.Equal(t, &ver, gotData)
			}
		})
	}
}

func TestVersionList(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *[]string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &[]string{"1.0", "0.9"},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
				Get(data_dragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Version.List()

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, &[]string{"1.0", "0.9"}, gotData)
			}
		})
	}
}

package cdragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/cdragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestVersionLatest(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := cdragon.NewCDragonClient(internalClient)

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
			wantErr: api.ErrNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.DataDragonURLFormat, "")).
				Get(cdragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Version.Latest()

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				ver := "1.0"
				require.Equal(t, ver, gotData)
			}
		})
	}
}

func TestVersionList(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := cdragon.NewCDragonClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    []string
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: []string{"1.0", "0.9"},
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
				Get(cdragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Version.List()

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				require.Equal(t, []string{"1.0", "0.9"}, gotData)
			}
		})
	}
}

package cdragon_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/cdragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestVersionLatest(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
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
			gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
				Get(cdragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			ctx := context.Background()
			gotData, gotErr := client.Version.Latest(ctx)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				ver := "1.0"
				require.Equal(t, ver, gotData)
			}
		})
	}
}

func TestVersionList(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
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
			gock.New(fmt.Sprintf(api.D_DRAGON_BASE_URL_FORMAT, "")).
				Get(cdragon.VersionsURL).
				Reply(test.code).
				JSON(test.want)

			ctx := context.Background()
			gotData, gotErr := client.Version.List(ctx)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				require.Equal(t, []string{"1.0", "0.9"}, gotData)
			}
		})
	}
}

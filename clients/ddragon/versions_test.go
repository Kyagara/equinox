package ddragon_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestVersionLatest(t *testing.T) {
	internal, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := ddragon.NewDDragonClient(internal)

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
				Get("/api/versions.json").
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
	internal, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := ddragon.NewDDragonClient(internal)

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
				Get("/api/versions.json").
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

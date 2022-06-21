package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestContentAllLocales(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.ContentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.ContentDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(val.ContentURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Content.AllLocales(val.BR)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

func TestContentByLocale(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.LocalizedContentDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.LocalizedContentDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(val.ContentURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Content.ByLocale(val.BR, val.PortugueseBR)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

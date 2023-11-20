package ddragon_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestRealmByName(t *testing.T) {
	internalClient, err := internal.NewInternalClient(equinox.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	client := ddragon.NewDDragonClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *ddragon.RealmData
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &ddragon.RealmData{},
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
				Get(fmt.Sprintf(ddragon.RealmURL, ddragon.BR)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Realm.ByName(ddragon.BR)
			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
			if test.wantErr == nil {
				require.Equal(t, test.want, gotData)
			}
		})
	}
}

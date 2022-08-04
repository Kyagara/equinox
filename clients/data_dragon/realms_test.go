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

func TestRealmByName(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := data_dragon.NewDataDragonClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *data_dragon.RealmData
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &data_dragon.RealmData{},
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
				Get(fmt.Sprintf(data_dragon.RealmURL, data_dragon.BR)).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Realm.ByName(data_dragon.BR)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

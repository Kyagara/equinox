package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlatformStatus(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *api.PlatformDataDTO
		wantErr error
		region  val.Shard
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &api.PlatformDataDTO{},
			region: val.BR,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.ErrNotFound,
			region:  val.BR,
		},
		{
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region ESPORTS is not available for this method"),
			region:  val.ESPORTS,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(val.StatusURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Status.PlatformStatus(test.region)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

package val_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/val"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestRankedLeaderboardsByActID(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := val.NewVALClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *val.LeaderboardDTO
		wantErr error
		region  val.Region
		size    uint8
		start   int
	}{
		{
			name:   "found",
			code:   http.StatusOK,
			want:   &val.LeaderboardDTO{},
			region: val.BR,
			size:   1,
			start:  0,
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
			region:  val.BR,
			size:    1,
			start:   0,
		},
		{
			name:    "invalid region",
			code:    http.StatusNotFound,
			wantErr: fmt.Errorf("the region ESPORTS is not available for this method"),
			region:  val.ESPORTS,
			size:    1,
			start:   0,
		},
		{
			name:   "default values",
			code:   http.StatusOK,
			want:   &val.LeaderboardDTO{},
			region: val.BR,
			size:   0,
			start:  -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, test.region)).
				Get(fmt.Sprintf(val.RankedURL, "actID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Ranked.LeaderboardsByActID(test.region, "actID", test.size, test.start)

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, test.want, gotData)
			}
		})
	}
}

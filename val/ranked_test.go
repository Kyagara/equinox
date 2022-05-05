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
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &val.LeaderboardDTO{},
		},
		{
			name:    "not found",
			code:    http.StatusNotFound,
			wantErr: api.NotFoundError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()

			gock.New(fmt.Sprintf(api.BaseURLFormat, val.BR)).
				Get(fmt.Sprintf(val.RankedURL, "actID")).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Ranked.LeaderboardsByActID(val.BR, "actID", 1, 0)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}
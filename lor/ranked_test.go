package lor_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestRankedLeaderboards(t *testing.T) {
	internalClient := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	client := lor.NewLORClient(internalClient)

	tests := []struct {
		name    string
		code    int
		want    *lor.LeaderboardDTO
		wantErr error
	}{
		{
			name: "found",
			code: http.StatusOK,
			want: &lor.LeaderboardDTO{},
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

			gock.New(fmt.Sprintf(api.BaseURLFormat, lor.Americas)).
				Get(lor.RankedURL).
				Reply(test.code).
				JSON(test.want)

			gotData, gotErr := client.Ranked.Leaderboards(lor.Americas)

			require.Equal(t, gotErr, test.wantErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))

			if test.wantErr == nil {
				assert.Equal(t, gotData, test.want)
			}
		})
	}
}

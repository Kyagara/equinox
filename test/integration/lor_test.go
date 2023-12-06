//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/require"
)

func TestLORPlatformStatus(t *testing.T) {
	checkIfOnlyDataDragon(t)
	ctx := context.Background()
	status, err := client.LOR.StatusV1.Platform(ctx, api.AMERICAS)
	require.NoError(t, err)
	require.NotEmpty(t, status, "expecting non-nil status")
	require.Equal(t, "Americas", status.Name, "expecting platform name to be equal to Americas")
}

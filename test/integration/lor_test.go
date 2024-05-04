//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/stretchr/testify/require"
)

func TestLORPlatformStatus(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	status, err := client.LOR.StatusV1.Platform(ctx, api.EUROPE)
	require.NoError(t, err)
	require.NotEmpty(t, status, "expecting non-nil status")
	require.Equal(t, "Europe", status.Name, "expecting platform name to be equal to Europe")
}

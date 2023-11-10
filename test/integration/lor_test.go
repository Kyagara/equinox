//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/require"
)

func TestLORPlatformStatus(t *testing.T) {
	checkIfOnlyDataDragon(t)
	status, err := client.LOR.StatusV1.Platform(api.AMERICAS)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, status, "expecting non-nil status")
}

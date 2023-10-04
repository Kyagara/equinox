//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/stretchr/testify/require"
)

func TestLORPlatformStatus(t *testing.T) {
	status, err := client.LOR.Status.PlatformStatus(lor.Americas)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, status, "expecting non-nil status")
}

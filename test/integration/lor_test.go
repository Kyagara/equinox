//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLORPlatformStatus(t *testing.T) {
	status, err := client.LOR.Status.PlatformStatus(lor.Americas)

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, status, "expecting non-nil status")
}

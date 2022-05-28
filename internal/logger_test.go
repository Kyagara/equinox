package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := internal.NewLogger(true, 10, 120, api.DebugLevel)

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

func TestLogger(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.NotNil(t, client, "expecting non-nil InternalClient")

	logger := client.Logger("logger", "s", "d")

	assert.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

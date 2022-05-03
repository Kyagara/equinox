package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := internal.NewLogger(internal.NewTestEquinoxConfig())

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

func TestLogger(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.NotNil(t, client, "expecting non-nil InternalClient")

	logger := client.Logger()

	assert.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

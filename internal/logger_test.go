package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger, err := internal.NewLogger(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, logger, "expecting non-nil Logger")
}

func TestLogger(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, internalClient, "expecting non-nil InternalClient")

	logger := internalClient.Logger("client", "endpoint", "method")

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

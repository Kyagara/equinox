package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	c := internal.NewTestEquinoxConfig()
	logger, err := internal.NewLogger(c)

	if err != nil {
		require.ErrorContains(t, err, "error initializing logger")
	}

	require.Equal(t, logger.Level(), zapcore.Level(-1))

	c.LogLevel = api.NopLevel

	logger, err = internal.NewLogger(c)

	if err != nil {
		require.ErrorContains(t, err, "error initializing logger")
	}

	assert.Equal(t, logger, zap.NewNop())
}

func TestLogger(t *testing.T) {
	internalClient, err := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, internalClient, "expecting non-nil InternalClient")

	logger := internalClient.Logger("client", "endpoint", "method")

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

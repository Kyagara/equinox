package internal_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	_, err := internal.NewLogger(nil)
	require.NotEmpty(t, err, "expecting non-nil error")

	config := &api.EquinoxConfig{
		LogLevel: api.NOP_LOG_LEVEL, Cache: &cache.Cache{TTL: 60 * time.Second},
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}

	logger, err := internal.NewLogger(config)
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, logger, "expecting non-nil logger")

	config.LogLevel = api.DEBUG_LOG_LEVEL
	logger, err = internal.NewLogger(config)
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, logger.Core().Enabled(zapcore.DebugLevel), true, "expecting logger to be enabled for Debug level")

	config.LogLevel = api.INFO_LOG_LEVEL
	logger, err = internal.NewLogger(config)
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, logger.Core().Enabled(zapcore.InfoLevel), true, "expecting logger to be enabled for Debug level")
}

func TestLogger(t *testing.T) {
	internal, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, internal, "expecting non-nil InternalClient")
	logger := internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger, "expecting non-nil Logger")
	logger = internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger, "expecting non-nil Logger")
	logger.Debug("Debug log")
	logger.Info("Info log")
	logger.Warn("Warn log")
	logger.Error("Error log")
}

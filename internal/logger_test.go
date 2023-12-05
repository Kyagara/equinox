package internal_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/test/util"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	config := &api.EquinoxConfig{
		LogLevel: zerolog.Disabled, Cache: &cache.Cache{TTL: 60 * time.Second},
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}

	logger := internal.NewLogger(config)
	require.NotEmpty(t, logger, "expecting non-nil logger")

	config.LogLevel = zerolog.DebugLevel
	logger = internal.NewLogger(config)
	require.Equal(t, logger.Debug().Enabled(), true, "expecting logger to be enabled for Debug level")

	config.LogLevel = zerolog.InfoLevel
	logger = internal.NewLogger(config)
	require.Equal(t, logger.Info().Enabled(), true, "expecting logger to be enabled for Debug level")
}

func TestLogger(t *testing.T) {
	internal, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, internal, "expecting non-nil InternalClient")
	logger := internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger, "expecting non-nil Logger")
	logger = internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger, "expecting non-nil Logger")
	logger.Debug().Msg("Debug log")
	logger.Info().Msg("Info log")
	logger.Warn().Msg("Warn log")
	logger.Error().Msg("Error log")
}

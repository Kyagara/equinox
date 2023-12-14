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
	logger := internal.NewLogger(api.EquinoxConfig{})
	require.Equal(t, zerolog.Disabled, logger.GetLevel())

	config := api.EquinoxConfig{
		Logger: api.Logger{
			Level: zerolog.Disabled,
		},
		Cache:      &cache.Cache{TTL: 60 * time.Second},
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}

	logger = internal.NewLogger(config)
	require.NotEmpty(t, logger)

	config.Logger.Level = zerolog.TraceLevel
	logger = internal.NewLogger(config)
	require.True(t, logger.Trace().Enabled())

	config.Logger.Level = zerolog.InfoLevel
	config.Logger.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = internal.NewLogger(config)
	require.True(t, logger.Info().Enabled())
}

func TestLogger(t *testing.T) {
	internal := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NotEmpty(t, internal)
	logger := internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger)
	logger = internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger)
	logger.Trace().Msg("Trace log")
	logger.Debug().Msg("Debug log")
	logger.Info().Msg("Info log")
	logger.Warn().Msg("Warn log")
	logger.Error().Msg("Error log")
}

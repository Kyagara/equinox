package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
	"github.com/Kyagara/equinox/v2/test/util"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := internal.NewLogger(api.EquinoxConfig{}, nil, nil)
	require.Equal(t, zerolog.Disabled, logger.GetLevel())

	config := api.EquinoxConfig{
		Logger: api.Logger{
			Level: zerolog.Disabled,
		},
	}

	logger = internal.NewLogger(config, nil, nil)
	require.NotEmpty(t, logger)

	config.Logger.Level = zerolog.TraceLevel
	logger = internal.NewLogger(config, nil, nil)
	require.True(t, logger.Trace().Enabled())

	config.Logger.Level = zerolog.InfoLevel
	config.Logger.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = internal.NewLogger(config, nil, nil)
	require.True(t, logger.Info().Enabled())
}

func TestLogging(t *testing.T) {
	internal := util.NewTestInternalClient(t)

	logger := internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger)

	// Logger already cached
	logger = internal.Logger("client_endpoint_method")
	require.NotEmpty(t, logger)

	logger.Trace().Msg("Trace log")
	logger.Debug().Msg("Debug log")
	logger.Info().Msg("Info log")
	logger.Warn().Msg("Warn log")
	logger.Error().Msg("Error log")
}

package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	logger := internal.NewLogger(config)

	require.NotNil(t, logger, "expecting non-nil Logger")
}

func TestLogger(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	require.NotNil(t, client, "expecting non-nil InternalClient")

	logger := client.Logger("client", "endpoint", "method")

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")
}

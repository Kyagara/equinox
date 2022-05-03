package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := internal.NewLogger(internal.NewTestEquinoxConfig())

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info("Info logger")

	logger.Warn("Warning logger")

	logger.Error("Error logger")
}

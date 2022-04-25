package internal_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	logger := internal.NewLogger()

	require.NotNil(t, logger, "expecting non-nil Logger")

	logger.Info.Println("Info logger")

	logger.Warn.Println("Warn logger")

	logger.Error.Println("Error logger")
}

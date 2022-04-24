package equinox_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := equinox.NewClient(os.Getenv("RIOT_API_KEY"))

	require.NotNil(t, client, "expecting non-nil client")
}

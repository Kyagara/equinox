package equinox_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/stretchr/testify/require"
)

func TestNewEquinoxClient(t *testing.T) {
	client := equinox.NewClientWithDebug(os.Getenv("RIOT_API_KEY"))

	require.NotNil(t, client, "expecting non-nil Equinox client")
}

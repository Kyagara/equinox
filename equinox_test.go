package equinox_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/stretchr/testify/assert"
)

func TestNewEquinoxClient(t *testing.T) {
	client, err := equinox.NewClientWithDebug(os.Getenv("RIOT_API_KEY"))

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}

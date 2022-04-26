package equinox_test

import (
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/assert"
)

func TestNewEquinoxClient(t *testing.T) {
	client, err := equinox.NewClientWithConfig(api.NewTestEquinoxConfig())

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}

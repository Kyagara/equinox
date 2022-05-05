package equinox_test

import (
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/assert"
)

func TestNewEquinoxClient(t *testing.T) {
	client, err := equinox.NewClient("RIOT_API_KEY")

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	config := &api.EquinoxConfig{
		Key:      "RIOT_API_KEY",
		LogLevel: api.DebugLevel,
		Timeout:  10,
		Retry:    true,
	}

	client, err := equinox.NewClientWithConfig(config)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}

func TestNewEquinoxClientWithoutKey(t *testing.T) {
	client, err := equinox.NewClient("")

	assert.Nil(t, client, "expecting nil Equinox client")

	assert.NotNil(t, err, "expecting non-nil error")
}

func TestNewEquinoxClientWithConfigWithoutKey(t *testing.T) {
	config := &api.EquinoxConfig{
		Key:      "",
		LogLevel: api.DebugLevel,
		Timeout:  10,
		Retry:    true,
	}

	client, err := equinox.NewClientWithConfig(config)

	assert.Nil(t, client, "expecting nil Equinox client")

	assert.NotNil(t, err, "expecting non-nil error")
}

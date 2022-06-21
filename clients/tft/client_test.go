package tft_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
)

func TestTFTClient(t *testing.T) {
	client := tft.NewTFTClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil TFTClient")
}

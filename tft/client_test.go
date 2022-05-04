package tft_test

import (
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/tft"
	"github.com/stretchr/testify/assert"
)

func TestTFTClient(t *testing.T) {
	client := tft.NewTFTClient(&internal.InternalClient{})

	assert.NotNil(t, client, "expecting non-nil TFTClient")
}

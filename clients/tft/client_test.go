package tft_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestTFTClient(t *testing.T) {
	client := tft.NewTFTClient(&internal.InternalClient{})

	require.NotNil(t, client, "expecting non-nil TFTClient")
}

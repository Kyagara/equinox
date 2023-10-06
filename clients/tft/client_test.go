package tft_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestTFTClient(t *testing.T) {
	c := &internal.InternalClient{}
	require.Equal(t, false, c.IsDataDragonOnly, "expecting non-nil TFTClient")
	client := tft.NewTFTClient(c)

	require.NotNil(t, client, "expecting non-nil TFTClient")

	c = &internal.InternalClient{IsDataDragonOnly: true}
	require.Equal(t, true, c.IsDataDragonOnly, "expecting non-nil TFTClient")
	client = tft.NewTFTClient(c)

	require.Nil(t, client, "expecting nil TFTClient")
}

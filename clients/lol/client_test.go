package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestLOLClient(t *testing.T) {
	client := lol.NewLOLClient(&internal.InternalClient{})

	require.NotNil(t, client, "expecting non-nil LOLClient")
}

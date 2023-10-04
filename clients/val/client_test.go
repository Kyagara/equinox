package val_test

import (
	"testing"

	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/require"
)

func TestVALClient(t *testing.T) {
	client := val.NewVALClient(&internal.InternalClient{})

	require.NotNil(t, client, "expecting non-nil VALClient")
}

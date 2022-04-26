package lol_test

import (
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestNewLOLClient(t *testing.T) {
	internalClient := internal.NewInternalClient(api.NewTestEquinoxConfig())

	client := lol.NewLOLClient(internalClient)

	assert.NotNil(t, client, "expecting non-nil LOLClient")
}

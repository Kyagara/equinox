package internal_test

import (
	"os"
	"testing"

	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
)

func TestNewInternalClient(t *testing.T) {
	internalClient := internal.NewInternalClient(os.Getenv("RIOT_API_KEY"), true)

	client := lol.NewLOLClient(internalClient)

	assert.NotNil(t, client, "expecting non-nil InternalClient")
}

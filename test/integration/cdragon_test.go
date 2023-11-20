//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCDragonChampionByName(t *testing.T) {
	version, err := client.CDragon.Version.Latest()
	require.Nil(t, err, "expecting nil error")
	champion, err := client.CDragon.Champion.ByName(version, "Lux")
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, champion, "expecting non-nil champion")
	require.Equal(t, "Lux", champion.Name, "expecting champion name to be Lux")
}

func TestCDragonChampionByID(t *testing.T) {
	version, err := client.CDragon.Version.Latest()
	require.Nil(t, err, "expecting nil error")
	champion, err := client.CDragon.Champion.ByID(version, 223)
	require.Nil(t, err, "expecting nil error")
	require.NotEmpty(t, champion, "expecting non-nil champion")
	require.Equal(t, "Tahm Kench", champion.Name, "expecting champion name to be Tahm Kench")
}

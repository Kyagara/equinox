//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/data_dragon"
	"github.com/stretchr/testify/require"
)

func TestDataDragonVersionLatest(t *testing.T) {
	version, err := client.DataDragon.Version.Latest()

	require.Nil(t, err, "expecting nil error")

	require.NotEqual(t, "", version, "expecting non-nil version")
}

func TestDataDragonRealmByName(t *testing.T) {
	realm, err := client.DataDragon.Realm.ByName(data_dragon.BR)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, realm, "expecting non-nil realm")

	require.Equal(t, "pt_BR", realm.L, "expecting realm language to be pt_BR")
}

func TestDataDragonChampionAllChampions(t *testing.T) {
	version, err := client.DataDragon.Version.Latest()

	require.Nil(t, err, "expecting nil error")

	champions, err := client.DataDragon.Champion.AllChampions(version, data_dragon.PtBR)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, champions, "expecting non-nil champions")

	require.Equal(t, "Jarvan IV", champions["JarvanIV"].Name, "expecting champion name to be Jarvan IV")
}

func TestDataDragonChampionByName(t *testing.T) {
	version, err := client.DataDragon.Version.Latest()

	require.Nil(t, err, "expecting nil error")

	champion, err := client.DataDragon.Champion.ByName(version, data_dragon.PtBR, "JarvanIV")

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, champion, "expecting non-nil champion")

	require.Equal(t, "Jarvan IV", champion.Name, "expecting champion name to be Jarvan IV")
}

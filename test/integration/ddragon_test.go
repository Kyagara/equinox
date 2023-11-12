//go:build integration
// +build integration

package integration

import (
	"testing"

	"github.com/Kyagara/equinox/clients/ddragon"
	"github.com/stretchr/testify/require"
)

func TestDDragonRealmByName(t *testing.T) {
	realm, err := client.DDragon.Realm.ByName(ddragon.BR)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, realm, "expecting non-nil realm")
	require.Equal(t, "pt_BR", realm.L, "expecting realm language to be pt_BR")
}

func TestDDragonChampionAllChampions(t *testing.T) {
	version, err := client.DDragon.Version.Latest()
	require.Nil(t, err, "expecting nil error")
	champions, err := client.DDragon.Champion.AllChampions(version, ddragon.PtBR)
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, champions, "expecting non-nil champions")
	require.Equal(t, true, len(champions) > 1, "expecting list to have more than one champions")
	require.Equal(t, "Jarvan IV", champions["JarvanIV"].Name, "expecting champion name to be Jarvan IV")
}

func TestDDragonChampionByName(t *testing.T) {
	version, err := client.DDragon.Version.Latest()
	require.Nil(t, err, "expecting nil error")
	champion, err := client.DDragon.Champion.ByName(version, ddragon.PtBR, "Lux")
	require.Nil(t, err, "expecting nil error")
	require.NotNil(t, champion, "expecting non-nil champion")
	require.Equal(t, "Lux", champion.Name, "expecting champion name to be Lux")
}

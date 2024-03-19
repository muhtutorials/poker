package p2p

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableAddPlayer(t *testing.T) {
	maxSeats := 2
	table := NewTable(maxSeats)

	assert.Nil(t, table.AddPlayer("1"))
	assert.Nil(t, table.AddPlayer("2"))

	assert.Equal(t, table.NumOfPlayers(), maxSeats)

	assert.NotNil(t, table.AddPlayer("3"))

	assert.Equal(t, table.NumOfPlayers(), maxSeats)
}

func TestTableGetPlayer(t *testing.T) {
	maxSeats := 10
	table := NewTable(maxSeats)

	for i := 0; i < maxSeats; i++ {
		addr := fmt.Sprintf("%d", i)
		assert.Nil(t, table.AddPlayer(addr))
		player, err := table.GetPlayer(addr)
		assert.Nil(t, err)
		assert.Equal(t, player.addr, addr)
	}
	assert.Equal(t, table.NumOfPlayers(), maxSeats)
}

func TestTableRemovePlayer(t *testing.T) {
	maxSeats := 10
	table := NewTable(maxSeats)

	for i := 0; i < maxSeats; i++ {
		addr := fmt.Sprintf("%d", i)
		assert.Nil(t, table.AddPlayer(addr))
		assert.Nil(t, table.RemovePlayer(addr))
		player, err := table.GetPlayer(addr)
		assert.NotNil(t, err)
		assert.Nil(t, player)
	}
	assert.Equal(t, table.NumOfPlayers(), 0)
}

func TestTableGetNextPlayer(t *testing.T) {
	maxSeats := 10
	table := NewTable(maxSeats)

	assert.Nil(t, table.AddPlayer("1"))
	assert.Nil(t, table.AddPlayer("2"))

	nextPlayer, err := table.GetNextPlayer("1")
	assert.Nil(t, err)
	assert.Equal(t, nextPlayer.addr, "2")

	assert.Nil(t, table.AddPlayer("3"))
	assert.Nil(t, table.RemovePlayer("2"))
	nextPlayer, err = table.GetNextPlayer("1")
	assert.Equal(t, nextPlayer.addr, "3")

	assert.Nil(t, table.RemovePlayer("3"))
	nextPlayer, err = table.GetNextPlayer("1")
	assert.Nil(t, nextPlayer)
	assert.NotNil(t, err)

	assert.Nil(t, table.AddPlayer("2"))
	nextPlayer, err = table.GetNextPlayer("2")
	assert.Nil(t, err)
	assert.Equal(t, nextPlayer.addr, "1")
}

func TestTableGetPrevPlayer(t *testing.T) {
	maxSeats := 10
	table := NewTable(maxSeats)

	assert.Nil(t, table.AddPlayer("1"))
	assert.Nil(t, table.AddPlayer("2"))

	prevPlayer, err := table.GetPrevPlayer("2")
	assert.Nil(t, err)
	assert.Equal(t, prevPlayer.addr, "1")

	prevPlayer, err = table.GetPrevPlayer("1")
	assert.Nil(t, err)
	assert.Equal(t, prevPlayer.addr, "2")
}

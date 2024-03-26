package p2p

import (
	"fmt"
	"strings"
	"sync"
)

type Player struct {
	addr     string
	tablePos int
	status   GameStatus
	action   PlayerAction
}

func NewPlayer(addr string, tablePos int) *Player {
	return &Player{
		addr:     addr,
		tablePos: tablePos,
		status:   AtTable,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s - %s", p.addr, p.status)
}

type Table struct {
	mu    sync.RWMutex
	seats map[int]*Player
	max   int
}

func NewTable(max int) *Table {
	return &Table{
		seats: make(map[int]*Player),
		max:   max,
	}
}

func (t *Table) String() string {
	var parts []string
	for _, p := range t.seats {
		format := fmt.Sprintf("[%d %s %s %s]", p.tablePos, p.addr, p.status, p.action)
		parts = append(parts, format)
	}
	return strings.Join(parts, " ")
}

func (t *Table) AddPlayer(addr string, pos int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.seats) == t.max {
		return fmt.Errorf("player table is full")
	}

	player := NewPlayer(addr, pos)

	t.seats[pos] = player

	return nil
}

func (t *Table) GetPlayer(addr string) (*Player, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.getPlayer(addr)
}

// without lock
func (t *Table) getPlayer(addr string) (*Player, error) {
	for i, _ := range t.seats {
		player, ok := t.seats[i]
		if ok {
			if player.addr == addr {
				return player, nil
			}
		}
	}

	return nil, fmt.Errorf("player (%s) not on the table", addr)
}

func (t *Table) GetNextPlayer(addr string) (*Player, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	player, err := t.GetPlayer(addr)
	if err != nil {
		return nil, err
	}

	nextPlayerPos := player.tablePos + 1

	for {
		nextPlayer, ok := t.seats[nextPlayerPos]
		if nextPlayer == player {
			return nil, fmt.Errorf("player (%s) is the only one on the table", addr)
		}
		if ok {
			return nextPlayer, nil
		}
		nextPlayerPos++
		if nextPlayerPos >= t.max {
			nextPlayerPos = 0
		}
	}
}

func (t *Table) GetPrevPlayer(addr string) (*Player, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	player, err := t.GetPlayer(addr)
	if err != nil {
		return nil, err
	}

	prevPlayerPos := player.tablePos - 1

	for {
		prevPlayer, ok := t.seats[prevPlayerPos]
		if prevPlayer == player {
			return nil, fmt.Errorf("player (%s) is the only one on the table", addr)
		}
		if ok {
			return prevPlayer, nil
		}
		prevPlayerPos--
		if prevPlayerPos <= 0 {
			prevPlayerPos = t.max
		}
	}
}

func (t *Table) RemovePlayer(addr string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i, _ := range t.seats {
		player, ok := t.seats[i]
		if ok {
			if player.addr == addr {
				delete(t.seats, i)
				return nil
			}
		}
	}

	return fmt.Errorf("player (%s) not on the table", addr)
}

func (t *Table) NumOfPlayers() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return len(t.seats)
}

func (t *Table) GetNextFreeSeat() int {
	for i := 0; i < t.max; i++ {
		if _, ok := t.seats[i]; !ok {
			return i
		}
	}
	panic("no free seat")
}

func (t *Table) Players() []*Player {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var players []*Player
	for _, player := range t.seats {
		players = append(players, player)
	}

	return players
}

func (t *Table) SetPlayerStatus(addr string, status GameStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()

	player, err := t.getPlayer(addr)
	if err != nil {
		panic(err)
	}
	player.status = status
}

func (t *Table) GetPlayerStatus(addr string) GameStatus {
	t.mu.RLock()
	defer t.mu.RUnlock()

	player, err := t.getPlayer(addr)
	if err != nil {
		panic(err)
	}
	return player.status
}

func (t *Table) SetPlayerAction(addr string, action PlayerAction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	player, err := t.getPlayer(addr)
	if err != nil {
		panic(err)
	}
	player.action = action
}

func (t *Table) NextRound(status GameStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i, _ := range t.seats {
		player, ok := t.seats[i]
		if ok {
			player.status = status
			player.action = PlayerActionNone
		}
	}
}

func (t *Table) EmptyTable() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.seats = make(map[int]*Player)
}

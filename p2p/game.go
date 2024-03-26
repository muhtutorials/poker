package p2p

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"poker/pb"
	"sync/atomic"
	"time"
)

type Game struct {
	addr         string
	broadcastCh  chan *Broadcast
	table        *Table
	gameStatus   GameStatus
	dealer       atomic.Int32
	playerList   *PlayerList
	playerTurn   atomic.Int32
	playerAction PlayerAction
}

func NewGame(addr string, bc chan *Broadcast) *Game {
	g := &Game{
		addr:        addr,
		broadcastCh: bc,
		table:       NewTable(6),
		playerList:  &PlayerList{},
	}
	g.playerList.add(addr)
	g.setPlayerTurn()
	go g.loop()
	return g
}

func (g *Game) loop() {
	for range time.Tick(time.Second * 5) {
		dealer, _ := g.getDealer()

		logrus.WithFields(logrus.Fields{
			"server":     g.addr,
			"playerList": g.playerList.list,
			"gameStatus": g.gameStatus,
			"dealer":     dealer,
			"playerTurn": g.playerTurn.Load(),
		}).Info("game state")
		logrus.WithFields(logrus.Fields{
			"table": g.table,
		}).Info("table state")
	}
}

func (g *Game) getDealer() (string, bool) {
	dealerAddr := g.playerList.get(int(g.dealer.Load()))
	return dealerAddr, g.addr == dealerAddr
}

func (g *Game) atomicSetGameStatus(s GameStatus) {
	atomic.StoreInt32((*int32)(&g.gameStatus), int32(s))
}

func (g *Game) atomicGetGameStatus() GameStatus {
	return GameStatus(atomic.LoadInt32((*int32)(&g.gameStatus)))
}

func (g *Game) atomicSetPlayerAction(a PlayerAction) {
	atomic.StoreInt32((*int32)(&g.playerAction), int32(a))
}

func (g *Game) atomicGetPlayerAction() PlayerAction {
	return PlayerAction(atomic.LoadInt32((*int32)(&g.playerAction)))
}

func (g *Game) AddPlayer(addr string) {
	g.playerList.add(addr)
}

func (g *Game) Broadcast(payload any, players ...string) {
	g.broadcastCh <- &Broadcast{
		To:      players,
		Payload: payload,
	}
}

// take seat after API call to "/take-seat" address
func (g *Game) takeSeatOut() {
	err := g.addPlayerToTable(g.addr)
	if err != nil {
		fmt.Println(err)
	}

	g.atomicSetGameStatus(AtTable)

	g.Broadcast(&pb.TakeSeatMsg{Addr: g.addr}, g.getOtherPlayers()...)
}

func (g *Game) takeSeatIn(addr string) {
	err := g.addPlayerToTable(addr)
	if err != nil {
		fmt.Println(err)
	}

	// if we don't have enough players the round can't be started
	if g.table.NumOfPlayers() < 4 {
		return
	}

	if _, isDealer := g.getDealer(); isDealer {
		go func() {
			fmt.Printf("I'm dealer (%s)\n", g.addr)
			if g.atomicGetGameStatus() == AtTable {
				g.shuffleAndEncryptOut(EncryptedDeck{})
			}
		}()
	}

	logrus.WithFields(logrus.Fields{
		"server": g.addr,
		"player": addr,
	}).Info("Setting player status to 'At table'")
}

func (g *Game) addPlayerToTable(addr string) error {
	tablePos := g.playerList.getIndex(addr)
	err := g.table.AddPlayer(addr, tablePos)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) shuffleAndEncryptOut(d EncryptedDeck) {
	nextPlayer, err := g.table.GetNextPlayer(g.addr)
	if err != nil {
		panic(err)
	}
	g.setGameStatusOut(ShuffleAndEncrypt)
	g.Broadcast(&pb.ShuffleAndEncryptMsg{Deck: d, Addr: g.addr}, nextPlayer.addr)
}

func (g *Game) shuffleAndEncryptIn(msg *pb.ShuffleAndEncryptMsg) error {
	prevPlayer, err := g.table.GetPrevPlayer(g.addr)
	if err != nil {
		panic(err)
	}

	if msg.Addr != prevPlayer.addr {
		return fmt.Errorf(
			"received encrypted deck from a wrong player (%s), should be (%s)",
			msg.Addr, prevPlayer)
	}

	_, isDealer := g.getDealer()

	if isDealer && msg.Addr == prevPlayer.addr {
		logrus.Info("Shuffle round complete")
		g.setGameStatusOut(PreFlop)
		return nil
	}

	g.shuffleAndEncryptOut(EncryptedDeck{})

	return nil
}

func (g *Game) getOtherPlayers() []string {
	var players []string
	for _, addr := range g.playerList.list {
		if addr == g.addr {
			continue
		}
		players = append(players, addr)
	}
	return players
}

func (g *Game) setGameStatusOut(s GameStatus) {
	if s != g.gameStatus {
		g.atomicSetGameStatus(s)
		g.table.SetPlayerStatus(g.addr, s)
		g.table.SetPlayerAction(g.addr, PlayerActionNone)
		g.Broadcast(&pb.SetGameStatusMsg{
			GameStatus: int32(s),
		}, g.getOtherPlayers()...)
	}
}

func (g *Game) setGameStatusIn(addr string, s GameStatus) {
	if g.isMessageFromDealer(addr) {
		g.atomicSetGameStatus(s)
		g.table.SetPlayerStatus(addr, s)
		g.table.SetPlayerAction(addr, PlayerActionNone)
		g.table.SetPlayerStatus(g.addr, s)
		g.table.SetPlayerAction(g.addr, PlayerActionNone)
		return
	}

	oldStatus := g.table.GetPlayerStatus(addr)
	if oldStatus == s {
		return
	}
	g.table.SetPlayerStatus(addr, s)
	g.Broadcast(&pb.SetGameStatusMsg{
		GameStatus: int32(s),
	}, g.getOtherPlayers()...)
}

func (g *Game) takeActionOut(action PlayerAction, value int) error {
	if !g.canTakeAction(g.addr) {
		return fmt.Errorf("player (%s) taking action before his turn", g.addr)
	}

	g.atomicSetPlayerAction(action)
	g.table.SetPlayerAction(g.addr, action)

	g.incrPlayerTurn()

	g.Broadcast(&pb.TakeActionMsg{
		Addr:         g.addr,
		GameStatus:   int32(g.gameStatus),
		PlayerAction: int32(action),
		Value:        int32(value),
	}, g.getOtherPlayers()...)

	if _, isDealer := g.getDealer(); isDealer {
		g.nextRound()
	}

	return nil
}

func (g *Game) takeActionIn(msg *pb.TakeActionMsg) error {
	if !g.canTakeAction(msg.Addr) {
		return fmt.Errorf("player (%s) taking action before his turn", msg.Addr)
	}

	if GameStatus(msg.GameStatus) != g.gameStatus && !g.isMessageFromDealer(msg.Addr) {
		return fmt.Errorf("player (%s) doesn't have correct game status", msg.Addr)
	}

	g.table.SetPlayerAction(msg.Addr, PlayerAction(msg.PlayerAction))

	g.incrPlayerTurn()

	if g.isMessageFromDealer(msg.Addr) {
		g.nextRound()
	}

	return nil
}

func (g *Game) canTakeAction(addr string) bool {
	playerTurnAddr := g.playerList.get(int(g.playerTurn.Load()))
	return playerTurnAddr == addr
}

func (g *Game) incrPlayerTurn() {
	if g.playerList.Len()-1 == int(g.playerTurn.Load()) {
		g.playerTurn.Store(0)
	} else {
		g.playerTurn.Add(1)
	}
}

func (g *Game) setPlayerTurn() {
	g.playerTurn.Store(g.dealer.Load() + 1)
}

func (g *Game) isMessageFromDealer(addr string) bool {
	return g.playerList.get(int(g.dealer.Load())) == addr
}

func (g *Game) nextRound() {
	g.atomicSetPlayerAction(PlayerActionNone)
	if g.atomicGetGameStatus() == River {
		//g.table.EmptyTable()
		g.takeSeatOut()
		return
	}
	nextGameStatus := g.getNextGameStatus()
	g.atomicSetGameStatus(nextGameStatus)
	g.table.NextRound(nextGameStatus)
}

func (g *Game) getNextGameStatus() GameStatus {
	switch g.gameStatus {
	case PreFlop:
		return Flop
	case Flop:
		return Turn
	case Turn:
		return River
	case River:
		return AtTable
	default:
		panic("invalid game status")
	}
}

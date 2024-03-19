package p2p

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

type Game struct {
	listenAddr   string
	broadcastCh  chan *Broadcast
	table        *Table
	status       GameStatus
	dealer       atomic.Int32
	playerList   *PlayerList
	playerTurn   atomic.Int32
	playerAction PlayerAction
}

func NewGame(addr string, bc chan *Broadcast) *Game {
	g := &Game{
		listenAddr:  addr,
		broadcastCh: bc,
		table:       NewTable(6),
		playerList:  &PlayerList{},
	}
	g.playerList.add(addr)
	go g.loop()
	return g
}

func (g *Game) loop() {
	for range time.Tick(time.Second * 5) {
		dealer, _ := g.getDealer()

		logrus.WithFields(logrus.Fields{
			"server":     g.listenAddr,
			"playerList": g.playerList.list,
			"gameStatus": g.status,
			"dealer":     dealer,
			"table":      g.table,
		}).Info("game state")
		logrus.WithFields(logrus.Fields{
			"table": g.table,
		}).Info("table state")
	}
}

func (g *Game) getDealer() (string, bool) {
	dealerAddr := g.playerList.get(int(g.dealer.Load()))
	return dealerAddr, g.listenAddr == dealerAddr
}

func (g *Game) setGameStatus(s GameStatus) {
	atomic.StoreInt32((*int32)(&g.status), int32(s))
}

func (g *Game) getGameStatus() GameStatus {
	return GameStatus(atomic.LoadInt32((*int32)(&g.status)))
}

func (g *Game) setPlayerAction(a PlayerAction) {
	atomic.StoreInt32((*int32)(&g.playerAction), int32(a))
}

func (g *Game) getPlayerAction() PlayerAction {
	return PlayerAction(atomic.LoadInt32((*int32)(&g.playerAction)))
}

func (g *Game) AddPlayer(from string) {
	g.playerList.add(from)
}

func (g *Game) sendToPlayers(payload any, players ...string) {
	g.broadcastCh <- &Broadcast{
		To:      players,
		Payload: payload,
	}
}

func (g *Game) SetReady() {
	tablePos := g.playerList.getIndex(g.listenAddr)
	g.table.AddPlayerOnPosition(g.listenAddr, tablePos)

	g.sendToPlayers(Ready, g.getOtherPlayers()...)
	g.SetStatus(Ready)
}

func (g *Game) deal() {
	fmt.Printf("I'm dealer (%s)\n", g.listenAddr)
	if g.getGameStatus() == Ready {
		g.InitShuffleAndDeal()
	}
}

func (g *Game) SetPlayerReady(from string) {
	tablePos := g.playerList.getIndex(from)
	g.table.AddPlayerOnPosition(from, tablePos)

	// if we don't have enough players the round can't be started
	if g.table.NumOfPlayers() < 3 {
		return
	}

	if _, areWeDealer := g.getDealer(); areWeDealer {
		go func() {
			g.deal()
		}()
	}

	logrus.WithFields(logrus.Fields{
		"we":     g.listenAddr,
		"player": from,
	}).Info("Setting player status to ready")
}

func (g *Game) InitShuffleAndDeal() {
	g.shuffleAndDeal(EncryptedDeck{})
}

func (g *Game) shuffleAndDeal(d EncryptedDeck) {
	dealToPlayer, err := g.table.GetNextPlayer(g.listenAddr)
	if err != nil {
		panic(err)
	}
	g.SetStatus(ShuffleAndDeal)
	g.sendToPlayers(EncryptedDeck{}, dealToPlayer.addr)
}

func (g *Game) getOtherPlayers() []string {
	var players []string
	for _, addr := range g.playerList.list {
		if addr == g.listenAddr {
			continue
		}
		players = append(players, addr)
	}
	return players
}

func (g *Game) SetStatus(s GameStatus) {
	if g.status != s {
		g.setGameStatus(s)
		g.table.SetPlayerStatus(g.listenAddr, s)
	}
}

func (g *Game) shuffleAndEncrypt(from string, deck EncryptedDeck) error {
	prevPlayer, err := g.table.GetPrevPlayer(g.listenAddr)
	if err != nil {
		panic(err)
	}

	if from != prevPlayer.addr {
		return fmt.Errorf(
			"received encrypted deck from a wrong player (%s), should be (%s)",
			from, prevPlayer)
	}

	_, isDealer := g.getDealer()

	if isDealer && from == prevPlayer.addr {
		logrus.Info("Shuffle round complete")
		g.SetStatus(PreFlop)
		g.table.SetPlayerStatus(g.listenAddr, PreFlop)
		g.sendToPlayers(PreFlop, g.getOtherPlayers()...)
		return nil
	}

	g.shuffleAndDeal(EncryptedDeck{})

	return nil
}

func (g *Game) TakeAction(action PlayerAction, value int) error {
	if !g.canTakeAction(g.listenAddr) {
		return fmt.Errorf("player (%s) taking action before his turn", g.listenAddr)
	}
	g.setPlayerAction(action)

	g.incrPlayerTurn()

	if g.playerList.get(int(g.dealer.Load())) == g.listenAddr {
		g.advanceToNextRound()
	}

	g.sendToPlayers(MessagePlayerAction{
		GameStatus:   g.status,
		PlayerAction: action,
		Value:        value,
	}, g.getOtherPlayers()...)

	return nil
}

func (g *Game) canTakeAction(from string) bool {
	playerAddr := g.playerList.get(int(g.playerTurn.Load()))

	return playerAddr == from
}

func (g *Game) incrPlayerTurn() {
	if g.playerList.Len()-1 == int(g.playerTurn.Load()) {
		g.playerTurn.Store(0)
	} else {
		g.playerTurn.Add(1)
	}
}

func (g *Game) handleOtherPlayerAction(msg MessagePlayerAction, from string) error {
	if !g.canTakeAction(from) {
		return fmt.Errorf("player (%s) taking action before his turn", from)
	}

	if msg.GameStatus != g.status && !g.isMessageFromDealer(from) {
		return fmt.Errorf("player (%s) doesn't have correct game status", from)
	}

	if g.playerList.get(int(g.dealer.Load())) == from {
		g.advanceToNextRound()
	}

	g.incrPlayerTurn()

	return nil
}

func (g *Game) isMessageFromDealer(from string) bool {
	return g.playerList.get(int(g.dealer.Load())) == from
}

func (g *Game) getNextGameStatus() GameStatus {
	switch g.status {
	case PreFlop:
		return Flop
	case Flop:
		return Turn
	case Turn:
		return River
	case River:
		return Ready
	default:
		panic("invalid game status")
	}
}

func (g *Game) advanceToNextRound() {
	g.setPlayerAction(PlayerActionNone)

	if g.getGameStatus() == River {
		g.SetReady()
		return
	}

	g.status = g.getNextGameStatus()
}

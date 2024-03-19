package p2p

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"sync"
)

const maxPlayers = 6

func init() {
	var gs GameStatus

	gob.Register(PeerList{})
	gob.Register(EncryptedDeck{})
	gob.Register(gs)
	gob.Register(MessagePlayerAction{})
}

const (
	TexasHoldem GameVariant = iota
	FiveCardStud
	FiveOPoker
	Guts
	Countdown
)

type GameVariant uint8

func (v GameVariant) String() string {
	switch v {
	case TexasHoldem:
		return "Texas Hold'em"
	case FiveCardStud:
		return "Five Card Stud"
	case FiveOPoker:
		return "Five-O Poker"
	case Guts:
		return "Guts"
	case Countdown:
		return "Countdown"
	default:
		return "Unknown game variant"
	}
}

type Config struct {
	ListenAddr    string
	APIListenAddr string
	GameVariant   GameVariant
	Version       string
	maxPlayers    int
}

type Server struct {
	Config
	game        *Game
	transport   *TCPTransport
	peersMu     sync.RWMutex
	peers       map[string]*Peer
	addPeerCh   chan *Peer
	delPeerCh   chan *Peer
	messageCh   chan *Message
	broadcastCh chan *Broadcast
}

func NewServer(cfg Config) *Server {
	if cfg.maxPlayers == 0 {
		cfg.maxPlayers = maxPlayers
	}

	s := &Server{
		Config: cfg,
		peers:  make(map[string]*Peer),
		// make it buffered in case of blocking
		addPeerCh:   make(chan *Peer, 10),
		delPeerCh:   make(chan *Peer, 10),
		messageCh:   make(chan *Message, 10),
		broadcastCh: make(chan *Broadcast, 10),
	}

	s.game = NewGame(s.ListenAddr, s.broadcastCh)
	s.transport = NewTCPTransport(cfg.ListenAddr, s.addPeerCh, s.delPeerCh)

	go func(s *Server) {
		apiServer := NewAPIServer(s.APIListenAddr, s.game)
		logrus.WithFields(logrus.Fields{
			"port": s.APIListenAddr,
		}).Info("API server started")
		apiServer.Run()
	}(s)

	return s
}

func (s *Server) Start() {
	go s.loop()

	logrus.WithFields(logrus.Fields{
		"port":       s.ListenAddr,
		"variant":    s.GameVariant,
		"version":    s.Version,
		"maxPlayers": s.maxPlayers,
	}).Info("Game server started")

	err := s.transport.Listen()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeerCh:
			s.handleAddPeer(peer)
		case peer := <-s.delPeerCh:
			s.handleDelPeer(peer)
		case msg := <-s.messageCh:
			go func() {
				if err := s.handleMessage(msg); err != nil {
					fmt.Println(err)
				}
			}()
		case msg := <-s.broadcastCh:
			go func() {
				if err := s.handleBroadcast(msg); err != nil {
					fmt.Println(err)
				}
			}()
		}
	}
}

func (s *Server) Connect(addr string) error {
	if s.isInPeerList(addr) {
		return nil
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	peer := &Peer{
		conn:       conn,
		listenAddr: s.ListenAddr,
		outbound:   true,
	}
	s.addPeerCh <- peer

	return s.sendHandshake(peer)
}

func (s *Server) isInPeerList(addr string) bool {
	for _, peer := range s.peers {
		if addr == peer.listenAddr {
			return true
		}
	}
	return false
}

func (s *Server) handleAddPeer(peer *Peer) {
	hs, err := s.receiveHandshake(peer)
	if err != nil {
		logrus.Error("Handshake with incoming player failed:", err)
		peer.conn.Close()
		s.handleDelPeer(peer)
		return
	}

	go peer.ReadLoop(s.messageCh)

	if !peer.outbound {
		if err := s.sendHandshake(peer); err != nil {
			logrus.Error("Failed to send handshake with peer:", err)
			peer.conn.Close()
			return
		}

		if err := s.sendPeerList(peer); err != nil {
			logrus.Error("Failed to send peer list to peer:", err)
			return
		}
	}

	s.addPeer(peer)

	s.game.AddPlayer(peer.listenAddr)

	logrus.WithFields(logrus.Fields{
		"variant":        hs.GameVariant,
		"version":        hs.Version,
		"peerRemoteAddr": peer.conn.RemoteAddr(),
		"peerListenAddr": peer.listenAddr,
		"we":             s.ListenAddr,
	}).Info("New peer connected")
}

func (s *Server) addPeer(peer *Peer) {
	s.peersMu.Lock()
	defer s.peersMu.Unlock()

	s.peers[peer.listenAddr] = peer
}

func (s *Server) handleDelPeer(peer *Peer) {
	delete(s.peers, peer.conn.RemoteAddr().String())
	logrus.WithFields(logrus.Fields{
		"addr": peer.conn.RemoteAddr(),
	}).Info("Peer disconnected")
}

func (s *Server) sendHandshake(peer *Peer) error {
	hs := &Handshake{
		GameVariant: s.GameVariant,
		Version:     s.Version,
		ListenAddr:  s.ListenAddr,
		GameStatus:  s.game.status,
	}

	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(hs); err != nil {
		return err
	}

	return peer.Send(buf.Bytes())
}

func (s *Server) receiveHandshake(peer *Peer) (*Handshake, error) {
	if len(s.peers) == s.maxPlayers {
		return nil, fmt.Errorf("max players exeeded (%d)", s.maxPlayers)
	}

	hs := &Handshake{}

	if err := gob.NewDecoder(peer.conn).Decode(hs); err != nil {
		return nil, err
	}

	if s.GameVariant != hs.GameVariant {
		return nil, fmt.Errorf("invalid game variant: %s", hs.GameVariant)
	}

	if s.Version != hs.Version {
		return nil, fmt.Errorf("wrong game variant: %s", hs.GameVariant)
	}

	peer.listenAddr = hs.ListenAddr

	return hs, nil
}

func (s *Server) sendPeerList(peer *Peer) error {
	peerList := s.getPeerList()

	if len(peerList) == 0 {
		return nil
	}

	message := &Message{
		From:    s.ListenAddr,
		Payload: peerList,
	}

	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(message); err != nil {
		return err
	}

	return peer.Send(buf.Bytes())
}

func (s *Server) getPeerList() PeerList {
	s.peersMu.RLock()
	defer s.peersMu.RUnlock()

	var peerList PeerList

	for _, p := range s.peers {
		peerList = append(peerList, p.listenAddr)
	}

	return peerList
}

func (s *Server) handleMessage(msg *Message) error {
	switch v := msg.Payload.(type) {
	case PeerList:
		s.handlePeerList(v)
	case EncryptedDeck:
		return s.handleEncryptedDeck(msg.From, v)
	case GameStatus:
		return s.handleGameStatus(v, msg.From)
	case MessagePlayerAction:
		return s.handleMessagePlayerAction(v, msg.From)
	}
	return nil
}

func (s *Server) handlePeerList(l PeerList) {
	for i := 0; i < len(l); i++ {
		if err := s.Connect(l[i]); err != nil {
			logrus.Error("failed to dial peer:", err)
			continue
		}
	}
}

func (s *Server) handleEncryptedDeck(from string, d EncryptedDeck) error {
	return s.game.shuffleAndEncrypt(from, d)
}

func (s *Server) handleBroadcast(b *Broadcast) error {
	msg := &Message{
		From:    s.ListenAddr,
		Payload: b.Payload,
	}

	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(msg); err != nil {
		return err
	}

	for _, addr := range b.To {
		peer, ok := s.peers[addr]
		if ok {
			go func(peer *Peer) {
				if err := peer.Send(buf.Bytes()); err != nil {
					logrus.Error("broadcast to peer error:", err)
				}
			}(peer)
		}
	}

	return nil
}

func (s *Server) handleGameStatus(gs GameStatus, from string) error {
	switch gs {
	case Ready:
		s.game.SetPlayerReady(from)
	case PreFlop:
		s.game.SetStatus(PreFlop)
	default:
		return fmt.Errorf("unknown game status")
	}

	return nil
}

func (s *Server) handleMessagePlayerAction(msg MessagePlayerAction, from string) error {
	logrus.WithFields(logrus.Fields{
		"from":   from,
		"we":     s.ListenAddr,
		"action": msg,
	}).Info("Received player action")
	return s.game.handleOtherPlayerAction(msg, from)
}

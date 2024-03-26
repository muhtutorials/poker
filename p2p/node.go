package p2p

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"poker/pb"
	"sync"
)

const maxPlayers = 6

type Config struct {
	Addr          string
	APIListenAddr string
	GameVariant   GameVariant
	Version       string
	MaxPlayers    int
}

type Node struct {
	pb.UnimplementedGossipServer
	Config
	peersLock   sync.RWMutex
	peers       map[string]pb.GossipClient
	broadcastCh chan *Broadcast
	game        *Game
}

func NewNode(config Config) *Node {
	if config.MaxPlayers == 0 {
		config.MaxPlayers = maxPlayers
	}

	n := &Node{
		Config:      config,
		peers:       make(map[string]pb.GossipClient),
		broadcastCh: make(chan *Broadcast, 10),
	}

	n.game = NewGame(n.Addr, n.broadcastCh)

	return n
}

func (n *Node) Start() error {
	ln, err := net.Listen("tcp", n.Addr)
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterGossipServer(gRPCServer, n)

	logrus.WithFields(logrus.Fields{
		"addr":       n.Addr,
		"variant":    n.GameVariant,
		"version":    n.Version,
		"maxPlayers": n.MaxPlayers,
	}).Info("Game server started")

	go func(n *Node) {
		apiServer := NewAPIServer(n.APIListenAddr, n.game)
		logrus.WithFields(logrus.Fields{
			"port": n.APIListenAddr,
		}).Info("API server started")
		apiServer.Run()
	}(n)

	go n.loop()

	return gRPCServer.Serve(ln)
}

func (n *Node) loop() {
	for bc := range n.broadcastCh {
		n.broadcast(bc)
	}
}

func (n *Node) Connect(addr string) error {
	if !n.canConnect(addr) {
		return nil
	}

	client, err := n.CreateClient(addr)
	if err != nil {
		return err
	}

	hs, err := client.ShakeHands(context.Background(), n.getHandshake())
	if err != nil {
		return err
	}

	n.addPeer(hs, client)

	return nil
}

func (n *Node) canConnect(addr string) bool {
	if addr == n.Addr {
		return false
	}

	for k := range n.peers {
		if k == addr {
			return false
		}
	}

	return true
}

func (n *Node) CreateClient(addr string) (pb.GossipClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewGossipClient(conn)

	return client, nil
}

func (n *Node) ShakeHands(ctx context.Context, hs *pb.Handshake) (*pb.Handshake, error) {
	client, err := n.CreateClient(hs.Addr)
	if err != nil {
		return nil, err
	}

	n.addPeer(hs, client)

	return n.getHandshake(), nil
}

func (n *Node) addPeer(hs *pb.Handshake, c pb.GossipClient) {
	n.peersLock.Lock()
	defer n.peersLock.Unlock()

	n.peers[hs.Addr] = c
	n.game.AddPlayer(hs.Addr)

	go func() {
		for _, addr := range hs.PeerAddrs {
			if err := n.Connect(addr); err != nil {
				continue
			}
		}
	}()
}

func (n *Node) getHandshake() *pb.Handshake {
	return &pb.Handshake{
		Version:     n.Version,
		GameVariant: uint32(n.GameVariant),
		GameStatus:  uint32(n.game.gameStatus),
		Addr:        n.Addr,
		PeerAddrs:   n.getPeerAddrs(),
	}
}

func (n *Node) broadcast(bc *Broadcast) {
	for _, addr := range bc.To {
		go func(addr string) {
			client, ok := n.peers[addr]
			if !ok {
				return
			}

			switch v := bc.Payload.(type) {
			case *pb.TakeSeatMsg:
				_, err := client.TakeSeat(context.Background(), v)
				if err != nil {
					fmt.Println(err)
				}
			case *pb.ShuffleAndEncryptMsg:
				_, err := client.ShuffleAndEncrypt(context.Background(), v)
				if err != nil {
					fmt.Println(err)
				}
			case *pb.SetGameStatusMsg:
				_, err := client.SetGameStatus(context.Background(), v)
				if err != nil {
					fmt.Println(err)
				} else {
					n.game.setGameStatusIn(addr, GameStatus(v.GameStatus))
				}
			case *pb.TakeActionMsg:
				_, err := client.TakeAction(context.Background(), v)
				if err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Println("unknown action")
			}
		}(addr)
	}
}

func (n *Node) getPeerAddrs() []string {
	n.peersLock.RLock()
	defer n.peersLock.RUnlock()

	addrs := make([]string, len(n.peers))
	for k := range n.peers {
		addrs = append(addrs, k)
	}

	return addrs
}

func (n *Node) TakeSeat(ctx context.Context, msg *pb.TakeSeatMsg) (*pb.Ack, error) {
	n.game.takeSeatIn(msg.Addr)

	return &pb.Ack{}, nil
}

func (n *Node) ShuffleAndEncrypt(ctx context.Context, msg *pb.ShuffleAndEncryptMsg) (*pb.Ack, error) {
	if err := n.game.shuffleAndEncryptIn(msg); err != nil {
		return nil, err
	}

	return &pb.Ack{}, nil
}

func (n *Node) SetGameStatus(ctx context.Context, msg *pb.SetGameStatusMsg) (*pb.Ack, error) {
	gameStatus := GameStatus(msg.GameStatus)
	n.game.setGameStatusIn(n.Addr, gameStatus)

	return &pb.Ack{}, nil
}

func (n *Node) TakeAction(ctx context.Context, msg *pb.TakeActionMsg) (*pb.Ack, error) {
	err := n.game.takeActionIn(msg)
	if err != nil {
		return nil, err
	}

	return &pb.Ack{}, nil
}

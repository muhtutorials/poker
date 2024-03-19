package p2p

import (
	"encoding/gob"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

type TCPTransport struct {
	listenAddr string
	addPeerCh  chan *Peer
	delPeerCh  chan *Peer
}

func NewTCPTransport(addr string, addPeerCh, delPeerCh chan *Peer) *TCPTransport {
	return &TCPTransport{
		listenAddr: addr,
		addPeerCh:  addPeerCh,
		delPeerCh:  delPeerCh,
	}
}

func (t *TCPTransport) Listen() error {
	lis, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		peer := &Peer{
			conn: conn,
		}
		t.addPeerCh <- peer
	}
}

type Peer struct {
	conn       net.Conn
	listenAddr string
	outbound   bool
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

func (p *Peer) ReadLoop(messageCh chan *Message) {
	for {
		message := new(Message)
		if err := gob.NewDecoder(p.conn).Decode(message); err != nil {
			fmt.Println(err)
			break
		}

		messageCh <- message
	}

	p.conn.Close()
}

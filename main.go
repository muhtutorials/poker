package main

import (
	"fmt"
	"net/http"
	"poker/p2p"
	"time"
)

func main() {
	player1 := createServerAndStart(":1000", ":1001")
	player2 := createServerAndStart(":2000", ":2001")
	player3 := createServerAndStart(":3000", ":3001")
	//player4 := createServerAndStart(":4000", ":4001")
	//player5 := createServerAndStart(":5000", ":5001")

	go sendReady(1, ":1001")
	go sendReady(2, ":2001")
	go sendReady(4, ":3001")
	//go sendReady(6, ":4001")
	//go sendReady(8, ":5001")

	player2.Connect(player1.ListenAddr)

	time.Sleep(time.Millisecond * 200)
	player3.Connect(player2.ListenAddr)

	//time.Sleep(time.Millisecond * 100)
	//player4.Connect(player3.ListenAddr)
	//
	//time.Sleep(time.Millisecond * 100)
	//player5.Connect(player4.ListenAddr)

	// flop
	time.Sleep(time.Second * 5)
	http.Get("http://localhost:1001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:2001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:3001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:4001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:5001/fold")
	//
	// turn
	time.Sleep(time.Second * 5)
	http.Get("http://localhost:1001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:2001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:3001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:4001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:5001/fold")
	//
	// river
	time.Sleep(time.Second * 5)
	http.Get("http://localhost:1001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:2001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:3001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:4001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:5001/fold")
	//
	// ready
	time.Sleep(time.Second * 5)
	http.Get("http://localhost:1001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:2001/fold")
	time.Sleep(time.Second)
	http.Get("http://localhost:3001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:4001/fold")
	//time.Sleep(time.Second)
	//http.Get("http://localhost:5001/fold")

	select {}
}

func createServerAndStart(addr, apiAddr string) *p2p.Server {
	config := p2p.Config{
		ListenAddr:    addr,
		APIListenAddr: apiAddr,
		GameVariant:   p2p.TexasHoldem,
		Version:       "Poker v0.1-alpha",
	}

	server := p2p.NewServer(config)

	go server.Start()

	time.Sleep(time.Millisecond * 100)

	return server
}

func sendReady(sleep int, addr string) {
	time.Sleep(time.Duration(sleep) * time.Second)
	http.Get(fmt.Sprintf("http://localhost%s/ready", addr))
}

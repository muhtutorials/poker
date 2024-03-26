package main

import (
	"fmt"
	"net/http"
	"poker/p2p"
	"time"
)

func main() {
	node1 := createServerAndStart(":1000", ":1001")
	node2 := createServerAndStart(":2000", ":2001")
	node3 := createServerAndStart(":3000", ":3001")
	node4 := createServerAndStart(":4000", ":4001")

	err := node2.Connect(node1.Addr)
	if err != nil {
		fmt.Println(err)
	}
	err = node3.Connect(node2.Addr)
	if err != nil {
		fmt.Println(err)
	}
	err = node4.Connect(node3.Addr)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 1)
	go takeSeat(1, node1.APIListenAddr)
	go takeSeat(2, node2.APIListenAddr)
	go takeSeat(3, node3.APIListenAddr)
	go takeSeat(4, node4.APIListenAddr)

	time.Sleep(5 * time.Second)

	// PREFLOP
	time.Sleep(time.Second * 2)
	http.Get("http://localhost:2001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:3001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:4001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:1001/fold")

	// FLOP
	time.Sleep(time.Second * 2)
	http.Get("http://localhost:2001/check")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:3001/check")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:4001/check")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:1001/check")

	// TURN
	time.Sleep(time.Second * 2)
	http.Get("http://localhost:2001/bet/2")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:3001/bet/3")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:4001/bet/4")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:1001/bet/1")

	// RIVER
	time.Sleep(time.Second * 2)
	http.Get("http://localhost:2001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:3001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:4001/fold")

	time.Sleep(time.Second * 2)
	http.Get("http://localhost:1001/fold")

	select {}
}

func createServerAndStart(addr, apiAddr string) *p2p.Node {
	config := p2p.Config{
		Addr:          addr,
		APIListenAddr: apiAddr,
		GameVariant:   p2p.TexasHoldem,
		Version:       "Poker v0.1-alpha",
	}

	node := p2p.NewNode(config)

	go func() {
		err := node.Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)

	return node
}

func takeSeat(sleep int, addr string) {
	time.Sleep(time.Duration(sleep) * time.Second)
	_, err := http.Get(fmt.Sprintf("http://localhost%s/take-seat", addr))
	if err != nil {
		fmt.Println(err)
	}
}

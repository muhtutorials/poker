package p2p

import (
	"sort"
	"sync"
)

type PlayerList struct {
	mu   sync.RWMutex
	list []string
}

func (p *PlayerList) add(addr string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.list = append(p.list, addr)
	sort.Sort(p)
}

func (p *PlayerList) get(index int) string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.list)-1 < index {
		panic("index out of range")
	}

	return p.list[index]
}

func (p *PlayerList) getIndex(listenAddr string) int {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, addr := range p.list {
		if listenAddr == addr {
			return i
		}
	}

	return -1
}

func (p *PlayerList) Len() int {
	return len(p.list)
}

func (p *PlayerList) Less(i, j int) bool {
	return p.list[i] < p.list[j]
}

func (p *PlayerList) Swap(i, j int) {
	p.list[i], p.list[j] = p.list[j], p.list[i]
}

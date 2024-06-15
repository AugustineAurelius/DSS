package node

import (
	"net"
	"sync"
)

type Peer struct {
	ID        [16]byte
	Secret    [32]byte
	RemotePub [65]byte
	con       net.Conn
	m         *sync.Mutex
}

const (
	Ping  = 12
	Pong  = 13
	IDReq = 14
)

func (p *Peer) lock() {
	p.m.Lock()
}
func (p *Peer) unlock() {
	p.m.Unlock()
}

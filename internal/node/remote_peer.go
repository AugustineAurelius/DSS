package node

import (
	"net"
	"sync"
)

type Peer struct {
	ID        [16]byte
	publicKey []byte
	m         *sync.Mutex
	con       net.Conn
}

func (p *Peer) lock() {
	p.m.Lock()
}
func (p *Peer) unlock() {
	p.m.Unlock()
}

func (p *Peer) Do(f func()) {
	p.lock()
	defer p.unlock()
	f()
}

package peer

import (
	"net"
	"sync"
)

type Remote struct {
	id        [16]byte
	publicKey []byte
	m         *sync.Mutex
	con       net.Conn
}

func New(c net.Conn) *Remote {
	return &Remote{
		con: c,
		m:   &sync.Mutex{},
	}

}

func (r *Remote) Do(f func()) {
	r.m.Lock()
	defer r.m.Unlock()
	f()
}

func (r *Remote) GetAddr() string {
	return r.con.RemoteAddr().String()
}

func (r *Remote) Conn() net.Conn {
	return r.con
}

func (r *Remote) ID() []byte {
	return r.id[:]
}

func (r *Remote) PublicKey() []byte {
	return r.publicKey
}

func (r *Remote) SetPublicKey(remPub []byte) {
	r.publicKey = remPub

}

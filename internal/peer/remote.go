package peer

import (
	"net"
	"sync"
	"sync/atomic"
)

type Remote struct {
	id        [16]byte
	publicKey [65]byte
	m         sync.Mutex
	con       net.Conn
	skip      atomic.Int32
}

func New(c net.Conn) *Remote {
	return &Remote{
		con:  c,
		m:    sync.Mutex{},
		skip: atomic.Int32{},
	}
}

func (r *Remote) IsSkip() bool {
	defer func() { r.skip.Store(0) }()
	return r.skip.Load() != 0
}

func (r *Remote) Skip() {
	r.skip.Store(1)
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
	return r.publicKey[:]
}

func (r *Remote) SetPublicKey(pk []byte) {
	copy(r.publicKey[:], pk)
}

func (r *Remote) IsEmpty() bool {
	for i := 0; i < len(r.publicKey); i++ {
		if r.publicKey[i] != 0 {
			return false
		}
	}

	return true
}

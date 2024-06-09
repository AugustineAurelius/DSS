package node

import "net"

type Peer struct {
	ID        [16]byte
	Secret    [32]byte
	RemotePub [65]byte
	con       net.Conn
}

const (
	Ping  = 12
	Pong  = 13
	IDReq = 14
)

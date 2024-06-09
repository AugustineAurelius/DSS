package node

type Peer struct {
	ID        [16]byte
	Secret    [32]byte
	RemotePub [65]byte
}

const (
	Ping  = 12
	Pong  = 13
	IDReq = 14
)

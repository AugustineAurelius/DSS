package node

import "crypto/ecdh"

type Peer struct {
	ID        [16]byte
	Secret    []byte
	RemotePub *ecdh.PublicKey
}

const (
	Ping = 12
	Pong = 13

	IDReq = 14
)

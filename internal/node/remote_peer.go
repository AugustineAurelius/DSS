package node

import "crypto/ecdh"

type Peer struct {
	ID        []byte
	Secret    []byte
	RemotePub *ecdh.PublicKey
}

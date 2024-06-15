package node

import (
	"net"
	"sync"

	"github.com/AugustineAurelius/DSS/pkg/buffer"
	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/message"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) handshake(c net.Conn) error {

	n.lock.Lock()
	defer n.lock.Unlock()

	peer := &Peer{con: c, m: &sync.Mutex{}}
	n.remotePeers = append(n.remotePeers, peer)

	err := n.keyExchange(peer.con)
	if err != nil {
		return err
	}

	// err = n.idExchange(c, peer)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (n *Node) keyExchange(c net.Conn) error {

	buf := buffer.Get()
	defer buffer.Put(buf)

	m := message.Get()
	defer message.Put(m)

	m.Type = message.KeyExchangeRequest
	codec.WriteHeader(m.Header[:], len(n.privateKey.PublicKey().Bytes()))
	m.Body = n.privateKey.PublicKey().Bytes()

	m.Encode(buf)

	err := write(c, buf)
	if err != nil {
		return err
	}

	return nil

}

func (n *Node) idExchange(c net.Conn, p *Peer) error {
	buf := buffer.Get()

	defer buffer.Put(buf)

	buf.Write(n.ID[:])
	write(c, buf)
	read(c, buf)

	resp := make([]byte, buf.Len())
	_, err := buf.Read(resp)
	if err != nil {
		return err
	}
	p.ID = [16]byte(resp)

	return nil
}

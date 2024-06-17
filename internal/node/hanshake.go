package node

import (
	"net"

	"github.com/AugustineAurelius/DSS/internal/peer"
	"github.com/AugustineAurelius/DSS/pkg/buffer"
	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/message"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) handshake(c net.Conn) error {

	n.lock.Lock()
	defer n.lock.Unlock()

	p := peer.New(c)
	n.remotePeers = append(n.remotePeers, p)

	err := n.keyExchange(p.Conn())
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) keyExchange(c net.Conn) error {

	buf := buffer.Get()
	defer buffer.Put(buf)

	m := message.Get()
	defer message.Put(m)

	m.Type = message.KeyExchangeRequest
	codec.WriteHeader(m.Header[:], 65)
	m.Body = n.privateKey.PublicKey().Bytes()

	m.Encode(buf)

	err := write(c, buf)
	if err != nil {
		return err
	}

	return nil

}

package node

import (
	"bytes"
	"errors"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) handshake(c net.Conn) error {

	peer := &Peer{}

	err := n.keyExchange(c, peer)
	if err != nil {
		return err
	}

	err = n.idExchange(c, peer)
	if err != nil {
		return err
	}

	n.lock.Lock()
	defer n.lock.Unlock()
	peer.con = c
	n.remoteNodes = append(n.remoteNodes, *peer)

	return nil
}

func (n *Node) keyExchange(c net.Conn, p *Peer) error {
	privateKey := ecdh.New()

	_, err := c.Write(privateKey.PublicKey().Bytes())
	if err != nil {
		return err
	}

	var pubBytes [65]byte
	_, err = c.Read(pubBytes[:])
	if err != nil {
		return err
	}
	p.RemotePub = pubBytes

	secret := ecdh.MustEDCH(privateKey, &pubBytes)

	_, err = c.Write(secret)
	if err != nil {
		return err
	}

	var remSecret [32]byte
	_, err = c.Read(remSecret[:])
	if err != nil {
		return err
	}
	p.Secret = remSecret

	if !bytes.Equal(secret, remSecret[:]) {
		c.Close()
		return errors.New("secret is not equal")
	}

	return nil

}

func (n *Node) idExchange(c net.Conn, p *Peer) error {
	var idReq [2]byte
	codec.Encode(&idReq, IDReq)

	_, err := c.Write(idReq[:])
	if err != nil {
		return err
	}

	_, err = c.Read(idReq[:])
	if err != nil {
		return err
	}

	if IDReq != codec.Decode(idReq[:]) {
		return errors.New("not id req")
	}

	_, err = c.Write(n.ID[:])
	if err != nil {
		return err
	}

	var remoteId [16]byte
	_, err = c.Read(remoteId[:])
	if err != nil {
		return err
	}

	p.ID = remoteId

	return nil
}

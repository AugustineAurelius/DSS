package node

import (
	"bytes"
	"errors"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) defaultECDHHandshake(c net.Conn) error {
	err := n.keyExchange(c)
	if err != nil {
		return err
	}

	err = n.idExchange(c)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) defaultECDHDial(c net.Conn) error {

	err := n.keyExchange(c)
	if err != nil {
		return err
	}
	err = n.idExchange(c)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node) keyExchange(c net.Conn) error {
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

	if !bytes.Equal(secret, remSecret[:]) {
		c.Close()
		return errors.New("secret is not equal")
	}

	return nil

}

func (n *Node) idExchange(c net.Conn) error {
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

	return nil
}

package node

import (
	"bytes"
	"errors"
	"fmt"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) defaultHandshake(c net.Conn) error {

	var b [2]byte
	codec.Ping(&b)
	_, err := c.Write(b[:])
	if err != nil {
		return err
	}

	_, err = c.Read(b[:])
	if err != nil {
		return err
	}

	fmt.Println("server get  pong", codec.Pong(b[:]))

	return nil
}

func (n *Node) defaultDial(c net.Conn) error {

	var b [2]byte

	_, err := c.Read(b[:])
	if err != nil {
		return err
	}

	codec.Ping(&b)
	_, err = c.Write(b[:])
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) defaultECDHHandshake(c net.Conn) error {

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

	dialerPub := ecdh.MustPublicKeyFromBytes(&pubBytes)

	secret, _ := privateKey.ECDH(dialerPub)

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

func (n *Node) defaultECDHDial(c net.Conn) error {

	privateKey := ecdh.New()

	var pubBytes [65]byte

	_, err := c.Read(pubBytes[:])
	if err != nil {
		return err
	}

	_, err = c.Write(privateKey.PublicKey().Bytes())
	if err != nil {
		return err
	}

	listenerPub := ecdh.MustPublicKeyFromBytes(&pubBytes)

	secret, _ := privateKey.ECDH(listenerPub)

	var remSecret [32]byte
	_, err = c.Read(remSecret[:])
	if err != nil {
		return err
	}

	if !bytes.Equal(secret, remSecret[:]) {
		c.Close()
		return errors.New("secret is not equal")
	}

	_, err = c.Write(secret)
	if err != nil {
		return err
	}

	return nil
}

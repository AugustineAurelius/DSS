package node

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/AugustineAurelius/DSS/pkg/buffer"
	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
)

type handshakeFunc = func(c net.Conn) error

func (n *Node) handshake(c net.Conn) error {

	peer := &Peer{}

	err := n.keyExchange(c)
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
	peer.m = &sync.Mutex{}
	n.remotePeers = append(n.remotePeers, *peer)

	return nil
}

func (n *Node) keyExchange(c net.Conn) error {
	privateKey := ecdh.New()

	buf := buffer.Get()
	defer buffer.Put(buf)

	buf.Write(privateKey.PublicKey().Bytes())
	write(c, buf)
	read(c, buf)

	remotePub := make([]byte, buf.Len())
	_, err := buf.Read(remotePub)
	if err != nil {
		return err
	}

	fmt.Println(remotePub)

	secret := ecdh.MustEDCH(privateKey, remotePub)
	buf.Write(secret)
	write(c, buf)
	read(c, buf)

	remoteSec := make([]byte, buf.Len())
	_, err = buf.Read(remoteSec)
	if err != nil {
		return err
	}

	if !bytes.Equal(secret, remoteSec) {
		c.Close()
		return errors.New("secret is not equal")
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

func read(c net.Conn, buf *bytes.Buffer) error {
	var header [2]byte
	_, err := c.Read(header[:])
	if err != nil {
		return err
	}

	body := make([]byte, codec.Decode(header[:]))

	_, err = c.Read(body)
	if err != nil {
		return err
	}

	_, err = buf.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func write(c net.Conn, buf *bytes.Buffer) {

	bufLen := buf.Len()

	body := make([]byte, bufLen)

	_, err := buf.Read(body)
	if err != nil {
		return
	}

	res := make([]byte, bufLen+2)

	codec.WriteHeader(res, uint16(bufLen))
	copy(res[2:], body)

	_, err = c.Write(res)
	if err != nil {
		return
	}
}

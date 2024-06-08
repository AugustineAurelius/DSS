package node

import (
	"fmt"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
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

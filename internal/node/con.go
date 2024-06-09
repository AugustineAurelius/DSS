package node

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/retry"
	"github.com/AugustineAurelius/DSS/pkg/uuid"
)

type Node struct {
	ID [16]byte

	lock sync.Mutex

	host string
	port string

	listener net.Listener

	remoteNodes []Peer
}

func New() *Node {

	return &Node{ID: uuid.New()}
}

func (n *Node) Serve() error {

	l, err := net.Listen(config.DefaultConfig.Network, n.port)
	if err != nil {
		return err
	}

	n.listener = l

	go n.acceptLoop()
	go n.pinger()
	return nil
}

func (n *Node) acceptLoop() {
	retry.Loop(func() error {
		conn, err := n.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			conn.Close()
			return err
		}

		err = n.handshake(conn)
		if err != nil {
			return err
		}
		go n.listen(conn)
		return nil

	}, time.Millisecond)
}

func (n *Node) dial(port string) error {

	conn, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return err
	}

	err = n.handshake(conn)
	if err != nil {
		return err
	}
	go n.listen(conn)

	return nil
}

func (n *Node) pinger() {
	retry.Loop(n.pingAll, time.Second)
}

func (n *Node) pingAll() error {
	for i := 0; i < len(n.remoteNodes); i++ {

		if err := n.pingOne(n.remoteNodes[i].con); err != nil {
			return err
		}
	}
	return nil
}
func (n *Node) pingOne(c net.Conn) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	var b [2]byte
	codec.Encode(&b, Ping)

	c.SetWriteDeadline(time.Now().Add(time.Second))
	_, err := c.Write(b[:])
	if err != nil {
		return err
	}

	return nil

}

func (n *Node) listen(c net.Conn) {

	f, err := retry.WrapForRetry(n.handle, c)
	if err != nil {
		return
	}
	retry.Loop(f, time.Millisecond*200)

}

func (n *Node) handle(c net.Conn) error {

	var b [2]byte

	c.SetReadDeadline(time.Now().Add(time.Second))

	_, err := c.Read(b[:])
	if err != nil {
		fmt.Println(err)
		return err

	}

	n.lock.Lock()
	defer n.lock.Unlock()

	c.SetWriteDeadline(time.Now().Add(time.Second))

	res := codec.Decode(b[:])

	switch res {
	case Ping:
		codec.Encode(&b, Pong)
		_, err = c.Write(b[:])
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("got ping", n.ID)
		return nil

	case Pong:
		fmt.Println("got pong", n.ID)
		return nil

	default:
		return errors.New("some err")
	}

}

func (n *Node) removePeer(index int) {

	n.lock.Lock()
	defer n.lock.Unlock()
	n.remoteNodes[index].con.Close()

	n.remoteNodes = append(n.remoteNodes[:index], n.remoteNodes[index+1:]...)

}

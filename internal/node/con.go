package node

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/codec"
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
	return nil
}

func (n *Node) acceptLoop() {
	for {

		conn, err := n.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			conn.Close()
			continue
		}

		err = n.defaultECDHHandshake(conn)
		if err != nil {
			return
		}

	}

}

func (n *Node) dial(port string) error {

	conn, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return err
	}

	err = n.defaultECDHDial(conn)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) pinger() {

	for {
		<-time.After(time.Second * 2)
		for i := 0; i < len(n.remoteNodes); i++ {
			err := n.ping(n.remoteNodes[i].con)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

}

func (n *Node) ponger(c net.Conn) {
	for {
		<-time.After(time.Second)

		var b [2]byte

		_, err := c.Read(b[:])
		if err != nil {
			fmt.Println(err)
			continue

		}

		res := codec.Decode(b[:])

		switch res {
		case Ping:
			codec.Encode(&b, Pong)
			_, err = c.Write(b[:])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("pong", n.ID)
		}
	}

}

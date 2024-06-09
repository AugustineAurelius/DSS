package node

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/uuid"
)

type Node struct {
	ID [16]byte

	lock sync.Mutex

	host string
	port string

	listener net.Listener

	remoteNodes []Peer
	connections []net.Conn
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

	time.Sleep(time.Second)
	return nil
}

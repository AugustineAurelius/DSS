package node

import (
	"fmt"
	"net"
	"sync"

	"github.com/AugustineAurelius/DSS/config"
)

type Node struct {
	ID []byte

	lock sync.Mutex

	host string
	port string

	listener net.Listener

	remoteNodes []Node
	connections []net.Conn
}

func New() *Node {

	return &Node{}
}

func (n *Node) Serve() error {

	l, err := net.Listen(config.DefaultConfig.Network, n.port)
	if err != nil {
		return err
	}

	n.listener = l

	//TODO start accept conns

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

		// var ping [2]byte
		// codec.Ping(ping)
		// conn.Write(ping[:])

	}

}

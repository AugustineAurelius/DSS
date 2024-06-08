package node

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/codec"
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

		fmt.Println("conn", conn)
		var b [2]byte

		codec.Ping(&b)
		_, err = conn.Write(b[:])
		if err != nil {
			fmt.Println(err)
			return
		}

		pong := make([]byte, 2)
		_, err = conn.Read(pong)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("server get  pong", codec.Pong(pong))

	}

}

func (n *Node) dial(port string) {

	con, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return
	}

	ping := make([]byte, 2)
	_, err = con.Read(ping)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("client got ping", codec.Pong(ping))

	var b [2]byte
	codec.Ping(&b)
	_, err = con.Write(b[:])
	if err != nil {
		fmt.Println(err)

		return
	}

	time.Sleep(time.Second)
}

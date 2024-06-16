package node

import (
	"fmt"
	"net"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/retry"
)

func (n *Node) Serve() error {

	l, err := net.Listen(config.DefaultConfig.Network, n.port)
	if err != nil {
		return err
	}

	n.listener = l

	go retry.Loop(n.accept, time.Microsecond)

	return nil
}

func (n *Node) accept() error {
	conn, err := n.listener.Accept()
	if err != nil {
		fmt.Printf("TCP accept error: %s\n", err)
		conn.Close()
		return err
	}

	fmt.Println("GOT new internal connection")

	err = n.handshake(conn)
	if err != nil {
		return err
	}

	return nil

}

package node

import (
	"net"
	"time"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/pkg/retry"
)

func (n *Node) dial(port string) error {

	conn, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return err
	}

	err = n.handshake(conn)
	if err != nil {
		return err
	}

	go retry.Loop(n.handle, time.Second)

	return nil
}

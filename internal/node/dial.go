package node

import (
	"net"

	"github.com/AugustineAurelius/DSS/config"
	"github.com/AugustineAurelius/DSS/internal/peer"
)

func (n *Node) Dial(port string) error {

	conn, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return err
	}

	n.lock.Lock()
	defer n.lock.Unlock()

	p := peer.New(conn)
	n.remotePeers = append(n.remotePeers, p)

	return nil
}

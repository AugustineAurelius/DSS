package node

import (
	"net"
	"sync"

	"github.com/AugustineAurelius/DSS/config"
)

func (n *Node) dial(port string) error {

	conn, err := net.Dial(config.DefaultConfig.Network, port)
	if err != nil {
		return err
	}

	n.lock.Lock()
	defer n.lock.Unlock()
	peer := &Peer{}
	peer.con = conn
	peer.m = &sync.Mutex{}
	n.remotePeers = append(n.remotePeers, peer)

	return nil
}

package node

import (
	"encoding/hex"
	"net"
	"sync"

	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/uuid"
)

type Node struct {
	ID [16]byte

	lock sync.Mutex

	host string
	port string

	listener    net.Listener
	remotePeers []Peer
}

func New() *Node {

	return &Node{ID: uuid.New()}
}

func (n *Node) handle() error {
	for i := 0; i < len(n.remotePeers); i++ {
		peer := n.remotePeers[i]

		peer.lock()

		if err := n.keyExchange(peer.con); err != nil {
			n.removePeer(i)
			return err
		}

		peer.unlock()
	}
	return nil
}

func (n *Node) removePeer(index int) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.remotePeers[index].con.Close()

	n.remotePeers = append(n.remotePeers[:index], n.remotePeers[index+1:]...)
}

func (n *Node) testSendHashes() {
	testHashes := "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"

	var we []byte
	var tm [2]byte

	hexByte := make([]byte, len(testHashes)/2)

	hex.Decode(hexByte, []byte(testHashes))

	codec.Encode(&tm, 128)
	we = append(we, tm[:]...)
	we = append(we, []byte(hexByte)...)

	n.remotePeers[0].con.Write(we)
}

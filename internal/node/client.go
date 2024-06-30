package node

import (
	"encoding/hex"
	"net"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/internal/peer"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
	"github.com/AugustineAurelius/DSS/pkg/retry"
	"github.com/AugustineAurelius/DSS/pkg/uuid"
	"github.com/syndtr/goleveldb/leveldb"
)

type Node struct {
	ID [16]byte

	privateKey *ecdh.PrivateKey

	lock sync.Mutex

	host string
	port string

	wg sync.WaitGroup

	listener    net.Listener
	remotePeers []*peer.Remote

	db *leveldb.DB
}

func New(port string) *Node {
	n := &Node{ID: uuid.New(), privateKey: ecdh.New(), port: port, wg: sync.WaitGroup{}, db: &leveldb.DB{}}

	db, err := leveldb.OpenFile("./storages/"+hex.EncodeToString(n.ID[:]), nil)
	if err != nil {
		panic(err)
	}
	n.db = db
	// defer n.db.Close()

	go retry.Loop(n.consume, time.Millisecond)

	return n
}

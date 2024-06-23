package node

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/AugustineAurelius/DSS/internal/peer"
	"github.com/AugustineAurelius/DSS/pkg/buffer"
	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/ecdh"
	"github.com/AugustineAurelius/DSS/pkg/message"
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

func (n *Node) consume() error {
	ch := make(chan net.Addr, len(n.remotePeers))

	for i := 0; i < len(n.remotePeers); i++ {
		if n.remotePeers[i].IsSkip() {
			continue
		}
		n.wg.Add(1)

		go func(index int) {
			defer n.wg.Done()
			n.readMsg(n.remotePeers[index], ch)
		}(i)
	}

	n.wg.Wait()

	close(ch)
	for v := range ch {
		n.removeByRemoteAddr(v)
	}

	return nil
}

func (n *Node) readMsg(peer *peer.Remote, ch chan<- net.Addr) {

	msg := message.Get()
	defer message.Put(msg)

	buf := buffer.Get()
	defer buffer.Put(buf)

	peer.Conn().SetReadDeadline(time.Now().Add(time.Second))

	err := read(peer.Conn(), buf)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			n.keyExchange(peer.Conn())
			return
		}
		ch <- peer.Conn().RemoteAddr()
		return
	}

	msg.Decode(buf)

	n.handleMessage(peer, msg, buf)
}

func (n *Node) handleMessage(peer *peer.Remote, msg *message.Payload, buf *bytes.Buffer) {
	switch msg.Type {
	case message.KeyExchangeRequest:
		peer.Do(
			func() {

				peer.SetPublicKey(msg.Body[:codec.Decode(msg.Header[:])])
				n.db.Put(msg.Hash(), msg.Body, nil)

				msg.Reset()

				msg.Type = message.KeyExchangeResponse
				codec.WriteHeader(msg.Header[:], 65)
				msg.Body = n.privateKey.PublicKey().Bytes()

				msg.Encode(buf)
				write(peer.Conn(), buf)

			},
		)

	case message.KeyExchangeResponse:
		peer.Do(
			func() {

				if peer.IsEmpty() {
					peer.SetPublicKey(msg.Body[:codec.Decode(msg.Header[:])])
				}
				n.db.Put(msg.Hash(), msg.Body, nil)

				msg.Reset()

				secret := ecdh.MustEDCH(n.privateKey, peer.PublicKey())

				msg.Type = message.SecretExchangeRequest
				codec.WriteHeader(msg.Header[:], len(secret))
				msg.Body = secret

				msg.Encode(buf)

				write(peer.Conn(), buf)

			},
		)

	case message.SecretExchangeRequest:
		peer.Do(
			func() {
				remoteSecret := msg.Body[:codec.Decode(msg.Header[:])]
				n.db.Put(msg.Hash(), msg.Body, nil)

				msg.Reset()

				secret := ecdh.MustEDCH(n.privateKey, peer.PublicKey())

				msg.Type = message.SecretExchangeResponse
				codec.WriteHeader(msg.Header[:], 1)

				if !bytes.Equal(secret, remoteSecret) {
					fmt.Println(errors.New("secret is not equal"), remoteSecret, secret)
					msg.Body = []byte{0}
				} else {
					msg.Body = []byte{1}
				}
				msg.Encode(buf)

				write(peer.Conn(), buf)
			},
		)

	case message.SecretExchangeResponse:
		if msg.Body[0] == 0 {
			n.removeByRemoteAddr(peer.Conn().RemoteAddr())
			return
		}
		n.db.Put(msg.Hash(), msg.Body, nil)

		msg.Reset()

		msg.Type = message.IDExchangeRequest
		codec.WriteHeader(msg.Header[:], 16)
		msg.Body = n.ID[:]

		msg.Encode(buf)

		write(peer.Conn(), buf)

	case message.IDExchangeRequest:
		peer.Do(
			func() {
				n.db.Put(msg.Hash(), msg.Body, nil)

				copy(peer.ID(), msg.Body)
				// msg.Reset()

				msg.Type = message.IDExchangeResponse
				codec.WriteHeader(msg.Header[:], 16)
				msg.Body = n.ID[:]

				msg.Encode(buf)

				write(peer.Conn(), buf)
			},
		)

	case message.IDExchangeResponse:
		peer.Do(
			func() {
				n.db.Put(msg.Hash(), msg.Body, nil)

				copy(peer.ID(), msg.Body)
				msg.Reset()

			},
		)
	default:

	}

}

func (n *Node) removePeer(index int) {
	n.lock.Lock()
	defer n.lock.Unlock()

	n.remotePeers[index].Conn().Close()
	n.remotePeers = append(n.remotePeers[:index], n.remotePeers[index+1:]...)
}

func (n *Node) removeByRemoteAddr(adr net.Addr) {
	for i := 0; i < len(n.remotePeers); i++ {

		if n.remotePeers[i].GetAddr() == adr.String() {
			n.removePeer(i)
			fmt.Println("removed", adr.String())
			return
		}
	}

}

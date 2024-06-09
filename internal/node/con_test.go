package node

import (
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {

	n1 := New()
	n1.port = ":4001"
	err := n1.Serve()
	if err != nil {
		t.Error(err)
		return
	}
	go n1.pinger()

	n2 := New()
	go n2.pinger()

	err = n2.dial(":4001")
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 10)

}

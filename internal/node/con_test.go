package node

import (
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {

	n1 := New()
	n1.n = "server"
	n1.port = ":4001"
	err := n1.Serve()
	if err != nil {
		t.Error(err)
		return
	}

	n2 := New()
	n2.n = "clietn"

	n2.port = ":4002"
	n2.Serve()

	err = n2.dial(":4001")
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 10)

}

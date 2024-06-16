package node

import (
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {

	n1 := New(":4001")
	err := n1.Serve()
	if err != nil {
		t.Error(err)
		return
	}

	n2 := New(":4002")

	n2.Serve()

	err = n2.Dial(":4001")
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep(time.Second * 10)

}

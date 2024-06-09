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

	n2 := New()
	n2.port = ":4002"
	n2.Serve()

	err = n2.dial(":4001")
	if err != nil {
		t.Error(err)
		return
	}

	// n3 := New()
	// err = n3.dial(":4001")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	//
	// err = n3.dial(":4002")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	time.Sleep(time.Second * 10)

}

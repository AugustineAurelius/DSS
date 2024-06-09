package node

import "testing"

func TestPingPong(t *testing.T) {

	n1 := Node{
		port: ":4001",
	}

	err := n1.Serve()
	if err != nil {
		t.Error(err)
		return
	}

	n2 := Node{}

	err = n2.dial(":4001")
	if err != nil {
		t.Error(err)
		return
	}

}

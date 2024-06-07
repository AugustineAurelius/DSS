package node

type Node struct {
	host string
	port string
}

func New() *Node {

	return &Node{}
}

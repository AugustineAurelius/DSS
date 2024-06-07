package node

import "net"

type handShakeFunc = func(c net.Conn) error

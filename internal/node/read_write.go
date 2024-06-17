package node

import (
	"bytes"
	"fmt"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
)

func read(c net.Conn, buf *bytes.Buffer) error {

	messageType := make([]byte, 1)

	n, err := c.Read(messageType[:])
	if err != nil || n == 0 {
		return fmt.Errorf("bad msg type %w", err)
	}

	_, err = buf.Write(messageType[:])
	if err != nil || n == 0 {
		return fmt.Errorf("bad msg type bw %w", err)
	}

	messageHeader := make([]byte, 2)
	_, err = c.Read(messageHeader)
	if err != nil || n == 0 {
		return err
	}

	_, err = buf.Write(messageHeader)
	if err != nil {
		return err
	}

	body := make([]byte, codec.Decode(messageHeader))
	_, err = c.Read(body)
	if err != nil {
		return err
	}

	_, err = buf.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func write(c net.Conn, buf *bytes.Buffer) error {

	m := make([]byte, buf.Len())

	_, err := buf.Read(m)
	if err != nil {
		return err
	}

	_, err = c.Write(m)
	if err != nil {
		return err
	}
	return nil
}

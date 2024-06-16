package node

import (
	"bytes"
	"net"

	"github.com/AugustineAurelius/DSS/pkg/codec"
)

func read(c net.Conn, buf *bytes.Buffer) error {

	messageType := make([]byte, 1)
	_, err := c.Read(messageType[:])
	if err != nil {
		return err
	}

	_, err = buf.Write(messageType[:])
	if err != nil {
		return err
	}

	var messageHeader [2]byte
	_, err = c.Read(messageHeader[:])
	if err != nil {
		return err
	}

	_, err = buf.Write(messageHeader[:])
	if err != nil {
		return err
	}

	body := make([]byte, codec.Decode(messageHeader[:]))

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

	_, err := buf.WriteTo(c)
	if err != nil {
		return err
	}
	return nil
}

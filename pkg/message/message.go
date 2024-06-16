package message

import (
	"bytes"

	"github.com/AugustineAurelius/DSS/pkg/codec"
)

type Payload struct {
	Type   byte
	Header [2]byte //2 bytes (body len)
	Body   []byte
}

func (p *Payload) Encode(buf *bytes.Buffer) {
	buf.WriteByte(p.Type)
	buf.Write(p.Header[:])
	buf.Write(p.Body)
}

func (p *Payload) Decode(buf *bytes.Buffer) {
	var err error
	p.Type, err = buf.ReadByte()
	if err != nil {
		panic(err)
	}

	buf.Read(p.Header[:])

	body := make([]byte, p.GetBodyLen())
	buf.Read(body)
	p.Body = body

}

func (p *Payload) GetBodyLen() uint16 {
	return codec.Decode(p.Header[:])
}

func (p *Payload) Reset() {
	p.Type = 0
	p.Header = [2]byte{0, 0}
	p.Body = p.Body[:0]
}

package message

import (
	"bytes"

	"github.com/AugustineAurelius/DSS/pkg/codec"
	"github.com/AugustineAurelius/DSS/pkg/crypto/hash"
)

type Payload struct {
	Type      byte
	Header    [2]byte //2 bytes (body len)
	PublicKey [32]byte
	Signature [64]byte
	Body      []byte
}

func (p *Payload) Encode(buf *bytes.Buffer) {
	buf.WriteByte(p.Type)
	buf.Write(p.Header[:])
	buf.Write(p.Body)
	p.Reset()
}

func (p *Payload) Decode(buf *bytes.Buffer) {
	t, err := buf.ReadByte()
	if err != nil {
		panic(err)
	}

	p.Type = t

	buf.Read(p.Header[:])

	body := make([]byte, p.GetBodyLen())

	buf.Read(body)
	p.Body = body
	buf.Reset()
}

func (p *Payload) GetBodyLen() uint16 {
	return codec.Decode(p.Header[:])
}

func (p *Payload) Reset() {
	p.Type = 0
	p.Header = [2]byte{0, 0}
	p.Body = p.Body[:0]
}

func (p *Payload) Hash() []byte {
	hash := hash.Hash512(p.Body)

	return hash[:]
}

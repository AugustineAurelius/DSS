package message

import (
	"sync"
)

var messagePool = sync.Pool{
	New: func() any { return new(Payload) },
}

func Get() *Payload {
	msg := messagePool.Get().(*Payload)
	if msg == nil {
		msg = new(Payload)
	}
	msg.Reset()
	return msg
}

func Put(p *Payload) {
	p.Reset()
	messagePool.Put(p)
}

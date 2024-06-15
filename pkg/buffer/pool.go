package buffer

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func Get() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	if buf == nil {
		buf = new(bytes.Buffer)
	}
	buf.Reset()
	return buf
}

func Put(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

package buffer

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any { return bytes.NewBuffer(make([]byte, 512)) },
}

func Get() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 512))
	}
	buf.Reset()
	return buf
}

func Put(buf *bytes.Buffer) {
	bufferPool.Put(buf)
}

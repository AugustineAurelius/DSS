package rlp

import (
	"testing"

	"github.com/AugustineAurelius/DSS/pkg/codec"
)

func TestEncodeLowByteArray(t *testing.T) {

	src := []byte("hello world!")
	dst := make([]byte, EncodeLen(src))

	Encode(dst, src)

	if codec.Decode(dst[0:2]) != 12 {
		t.Error("bad lenght")
		return
	}
}

func TestEncodeLongByteArray(t *testing.T) {

	src := []byte("hello world!hello world!hello world!hello world!hello world!")
	dst := make([]byte, EncodeLen(src))

	Encode(dst, src)

	if codec.Decode(dst[0:2]) != 60 {
		t.Error("bad lenght")
		return
	}
}
func BenchmarkEncodeLowByteArray(b *testing.B) {
	src := []byte("hello world!")
	dst := make([]byte, len(src)+2)

	for i := 0; i < b.N; i++ {
		Encode(dst, src)
	}
}

func BenchmarkEncodeLongByteArray(b *testing.B) {
	src := []byte("hello world!hello world!hello world!hello world!hello world!")
	dst := make([]byte, len(src)+2)

	for i := 0; i < b.N; i++ {
		Encode(dst, src)
	}
}

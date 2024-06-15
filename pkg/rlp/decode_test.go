package rlp

import (
	"testing"
)

func TestDecode(t *testing.T) {
	str := []byte("hi hih ih ih ihihihihihi")
	enc := make([]byte, EncodeLen(str))

	Encode(enc, str)

	dst := make([]byte, DecodeLen(enc))
	Decode(dst, enc)
}

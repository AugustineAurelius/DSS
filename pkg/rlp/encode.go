package rlp

import (
	"github.com/AugustineAurelius/DSS/pkg/codec"
)

func Encode(dst, src []byte) {
	var buf [2]byte
	codec.Encode(&buf, uint16(len(src)))

	copy(dst[0:2], buf[:])
	copy(dst[2:], src)
}

func EncodeLen(src []byte) int {
	return len(src) + 2
}

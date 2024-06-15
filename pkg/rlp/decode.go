package rlp

func Decode(dst, src []byte) {
	copy(dst, src[2:])
}

func DecodeLen(src []byte) int {
	return len(src) - 2
}

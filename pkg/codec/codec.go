package codec

//little endian

func Encode(b *[2]byte, req uint16) {
	b[0] = byte(req)
	b[1] = byte(req >> 8)
}

func Decode(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

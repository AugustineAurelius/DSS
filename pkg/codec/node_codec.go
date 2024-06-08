package codec

//little endian

const ping = 12

func Ping(b *[2]byte) {
	b[0] = byte(ping)
	b[1] = byte(ping >> 8)
}

func Pong(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

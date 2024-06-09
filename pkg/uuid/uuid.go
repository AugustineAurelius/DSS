package uuid

import (
	"crypto/rand"
)

func New() [16]byte {
	var buf [16]byte

	_, err := rand.Read(buf[:])
	if err != nil {
		panic(err)
	}

	return buf
}

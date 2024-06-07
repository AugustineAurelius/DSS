package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/hex"
)

var curve = ecdh.P256()

func NewPair(private *[64]byte, public *[130]byte) {
	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	hex.Encode(private[:], privateKey.Bytes())

	hex.Encode(public[:], privateKey.PublicKey().Bytes())
}

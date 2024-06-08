package ecdh

import (
	"crypto/ecdh"
	"crypto/rand"
)

var curve = ecdh.P256()

// public len = 65; private len = 32
func New() *ecdh.PrivateKey {
	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return privateKey

}

func MustPrivateKeyFromBytes(key *[32]byte) *ecdh.PrivateKey {

	privateKey, err := curve.NewPrivateKey(key[:])
	if err != nil {
		panic(err)
	}
	return privateKey
}
func MustPublicKeyFromBytes(key *[65]byte) *ecdh.PublicKey {

	publicKey, err := curve.NewPublicKey(key[:])
	if err != nil {
		panic(err)

	}

	return publicKey

}
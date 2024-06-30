package ed25519

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"io"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
	// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
	SeedSize = 32
)

var (
	opts     = &ed25519.Options{Hash: crypto.SHA512}
	hashFunc = sha512.Sum512
)

func New() (ed25519.PublicKey, ed25519.PrivateKey) {

	seed := make([]byte, SeedSize)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		panic(err)
	}

	privateKey := ed25519.NewKeyFromSeed(seed)

	publicKey := make([]byte, PublicKeySize)
	copy(publicKey, privateKey[32:])

	return publicKey, privateKey
}

func MustSign(private ed25519.PrivateKey, msg []byte) []byte {
	hashOfMsg := hashFunc(msg)

	signature, err := private.Sign(rand.Reader, hashOfMsg[:], opts)
	if err != nil {
		panic(err)
	}

	return signature
}

func Verify(public ed25519.PublicKey, msg, signature []byte) bool {
	hashOfMsg := hashFunc(msg)

	return nil == ed25519.VerifyWithOptions(public, hashOfMsg[:], signature, opts)
}

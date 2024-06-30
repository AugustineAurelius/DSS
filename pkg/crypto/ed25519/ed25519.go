package ed25519

import (
	"crypto"
	"crypto/ed25519"

	hashbuffer "github.com/AugustineAurelius/DSS/pkg/crypto/hash_buffer"

	"crypto/rand"
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
	opts = &ed25519.Options{Hash: crypto.SHA512}
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

	hash := hashbuffer.Default.Hash(msg)

	signature, err := private.Sign(rand.Reader, hash[:], opts)
	if err != nil {
		panic(err)
	}

	return signature
}

// encoded signature should be equal to public key
func Verify(public, signature []byte) bool {
	hash := hashbuffer.Default.Hash(public)
	return nil == ed25519.VerifyWithOptions(public, hash[:], signature, opts)
}

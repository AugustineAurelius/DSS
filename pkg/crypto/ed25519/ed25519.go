package ed25519

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha512"

	"crypto/rand"
	"io"

	"github.com/AugustineAurelius/DSS/pkg/crypto/hash"
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

type PrivateKey struct{ ed25519.PrivateKey }

type PublicKey = ed25519.PublicKey

func (pk PrivateKey) PublicKey() []byte {
	return pk.PrivateKey[32:]
}

func (pk PrivateKey) Sign(p []byte) [64]byte {
	hash := hash.Hash512(p)

	signature, err := pk.PrivateKey.Sign(rand.Reader, hash[:], opts)
	if err != nil {
		panic(err)
	}

	return [64]byte(signature)
}

var (
	opts = &ed25519.Options{Hash: crypto.SHA512}
)

func New() (PublicKey, PrivateKey) {

	seed := make([]byte, SeedSize)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		panic(err)
	}

	privateKey := ed25519.NewKeyFromSeed(seed)

	publicKey := make([]byte, PublicKeySize)
	copy(publicKey, privateKey[32:])

	return publicKey, PrivateKey{privateKey}
}

func MustSign(private PrivateKey, msg []byte) []byte {

	signature := private.Sign(msg)

	return signature[:]
}

// encoded signature should be equal to public key
func Verify(public, signature []byte) bool {
	hash := sha512.Sum512(public)
	return nil == ed25519.VerifyWithOptions(public, hash[:], signature, opts)
}

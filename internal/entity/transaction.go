package entity

import (
	"bytes"
	"math/big"

	"github.com/AugustineAurelius/DSS/pkg/crypto/ed25519"
)

type Transaction struct {
	Timestamp int64

	Block uint32

	Tx *SignedTransaction

	GasUsed uint16
}

type RawTransaction struct {
	Sender   [32]byte //32
	Receiver [32]byte //32

	Amount *big.Int

	GasPrice uint16
	GasLimit uint16

	Nonce uint64

	Data []byte
}

func (r *RawTransaction) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.Grow(32 + 32 + r.Amount.BitLen() + len(r.Data))

	buf.Write(r.Sender[:])
	buf.Write(r.Receiver[:])

	buf.Write(r.Amount.Bytes())
	buf.Write(r.Data)

	return buf.Bytes()

}

type SignedTransaction struct {
	Tx *RawTransaction

	publicKey [ed25519.PublicKeySize]byte
	Signature [ed25519.SignatureSize]byte //sender public address is hashed by his private key
}

// TODO Signer interface
func Sign(tx *RawTransaction, signer ed25519.PrivateKey) SignedTransaction {
	return SignedTransaction{
		Tx:        tx,
		publicKey: [32]byte(signer.PublicKey()),
		Signature: signer.Sign(tx.Bytes()),
	}
}

func (t *Transaction) Reset() {
	*t = Transaction{}
}

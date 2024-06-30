package entity

import (
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
	Sender   []byte //32
	Receiver []byte //32

	Amount *big.Int

	GasPrice uint16
	GasLimit uint16

	Nonce uint64

	Data []byte
}

type SignedTransaction struct {
	Tx *RawTransaction

	publicKey [ed25519.PublicKeySize]byte
	Signature [ed25519.SignatureSize]byte //sender public address is hashed by his private key
}

func (t *Transaction) Reset() {
	*t = Transaction{}
}

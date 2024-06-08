package entity

import "math/big"

type Transaction struct {
	Timestamp int64

	Block uint32

	Tx *SignedTransaction

	GasUsed uint16
}

type RawTransaction struct {
	Sender   []byte
	Receiver []byte

	Amount *big.Int

	GasPrice uint16
	GasLimit uint16

	Nonce uint64

	Data []byte
}

type SignedTransaction struct {
	Tx *RawTransaction

	Signature []byte //sender public address is hashed by his private key
}

func (t *Transaction) Reset() {
	*t = Transaction{}
}

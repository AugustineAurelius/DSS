package node

import (
	"math/big"

	"github.com/AugustineAurelius/DSS/internal/entity"
)

func (n *Node) sendTransacrion() {

	raw := entity.RawTransaction{}

	raw.Sender = [32]byte{1}
	raw.Receiver = [32]byte{1}
	raw.Amount = big.NewInt(123)

	raw.GasPrice = 10
	raw.GasLimit = 1
	raw.Nonce = 0

}

func testTransaction() entity.RawTransaction {
	return entity.RawTransaction{
		Sender:   [32]byte{1, 1, 1},
		Receiver: [32]byte{1, 1, 1, 1},
		Amount:   big.NewInt(123123),
		GasPrice: 10,
		GasLimit: 123,
	}
}

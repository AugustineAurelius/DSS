package entity

type Block struct {
	Height    uint32
	Timestamp int64

	Transactions []Transaction

	BlockHash  []byte
	RootHash   []byte
	ParentHash []byte

	fee  uint32
	size uint32
}

func (b *Block) Reset() {
	*b = Block{}
}

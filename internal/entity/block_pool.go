package entity

import "sync"

type BlockPool struct {
	basePool sync.Pool
}

func NewBlockPool() *BlockPool {
	return &BlockPool{
		basePool: sync.Pool{New: func() any {
			return new(Block)
		}},
	}

}

func (p *BlockPool) Get() *Block {
	block := p.basePool.Get()
	if block == nil {
		return new(Block)
	}

	return block.(*Block)
}

func (p *BlockPool) Return(block *Block) {
	p.basePool.Put(block)
}

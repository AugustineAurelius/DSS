package entity

import "sync"

type TransactionPool struct {
	basePool sync.Pool
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		basePool: sync.Pool{New: func() any {
			return new(Transaction)
		}},
	}

}

func (p *TransactionPool) Get() *Transaction {
	transaction := p.basePool.Get()
	if transaction == nil {
		return new(Transaction)
	}

	return transaction.(*Transaction)
}

func (p *TransactionPool) Return(transaction *Transaction) {
	p.basePool.Put(transaction)
}

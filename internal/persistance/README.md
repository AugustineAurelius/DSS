unique timestamp 	time.Now().UTC().UnixNano()


gas base -> const


Block{
	blocknumber
	timestamp
	block hash
	[] transaction
	count of tx

	hash
	merkletree hash

	hash of prev block
	fee
	size

}


userfriendly
transaction {
	block *Block
	timestamp
	sender
	receiver
	amount
	gas price uint16
	gas limit uint16
	gas used uint16
	fee float32
	nonce uint64
	data []byte
}


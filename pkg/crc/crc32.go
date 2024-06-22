package crc

import "hash/crc32"

func CheckSum(msg []byte) uint64 {

	crc32.New(crc32.MakeTable(crc32.Koopman))
	return 0
}

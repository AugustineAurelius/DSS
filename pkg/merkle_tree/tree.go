package merkletree

import (
	"bytes"
	"crypto/sha256"
)

func CalculateHash256(dst *[32]byte, src []byte) {
	*dst = sha256.Sum256(src)
}

func validate(bufCap int) bool {
	if bufCap%32 != 0 {
		return false
	}
	return true
}

// used sha256
func MerkleTree32(dst *[32]byte, buf *bytes.Buffer) {
	cp := buf.Cap()
	if !validate(cp) {
		return
	}

	mp := make([][32]byte, 0, 32)

	for buf.Len() != 0 {

		temp := make([]byte, 64)
		readedBytes, err := buf.Read(temp)
		if err != nil {
			return
		}

		if readedBytes == 32 {
			temp = append(temp[:32], temp[:32]...)
		}
		var res [32]byte
		CalculateHash256(&res, temp)

		mp = append(mp, res)
	}

	for len(mp) != 1 {
		l := len(mp)
		intermedaiteMap := make([][32]byte, 0, l/2)

		for i := 0; i < l; i = i + 2 {
			var tmp [32]byte
			temp := make([]byte, 0, 64)

			temp = append(temp, mp[i][:]...)

			if i+1 == len(mp) {
				i--
			}

			temp = append(temp, mp[i+1][:]...)

			CalculateHash256(&tmp, temp)

			intermedaiteMap = append(intermedaiteMap, tmp)
		}

		mp = intermedaiteMap

	}
	*dst = mp[0]
}

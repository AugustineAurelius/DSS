package merkletree

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestTree2(t *testing.T) {

	str := "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"

	hexByte := make([]byte, len(str)/2)

	hex.Decode(hexByte, []byte(str))

	buf := bytes.NewBuffer(hexByte)

	var res [32]byte
	MerkleTree32(&res, buf)

	dst := make([]byte, 64)
	hex.Encode(dst, res[:])

	fmt.Println(string(dst))

}

func BenchmarkTreeEncode(b *testing.B) {

	str := "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"

	hexByte := make([]byte, len(str)/2)

	hex.Decode(hexByte, []byte(str))

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(hexByte)
		var res [32]byte
		MerkleTree32(&res, buf)

		dst := make([]byte, 64)
		hex.Encode(dst, res[:])

	}

}

func BenchmarkTreeEncodeString(b *testing.B) {

	str := "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"

	hexByte := make([]byte, len(str)/2)

	hex.Decode(hexByte, []byte(str))

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(hexByte)
		var res [32]byte
		MerkleTree32(&res, buf)

		hex.EncodeToString(res[:])

	}

}

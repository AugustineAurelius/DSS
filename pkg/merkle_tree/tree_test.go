package merkletree

import (
	"bytes"
	"encoding/hex"
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

	expected := "964348c9a9891b3d04b32541517e113a181fc70dd1c0b89858539192790b6a43"

	if string(dst) != expected {
		t.Error("result is not equal with expected")
	}

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

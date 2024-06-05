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

//	func TestTreeEvenContent(t *testing.T) {
//		opts := MerkleTreeOpts{
//			HashFunc: DefaultHashFunc,
//		}
//
//		contents := make([]Content, 0, 2)
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//
//		tree, err := NewMerkleTree(opts, contents)
//		if err != nil {
//			t.Error("construct err", err)
//		}
//
//		byteHash, err := tree.CalculateHash()
//		if err != nil {
//			t.Error(err)
//		}
//
//		expectedHash := "5924f0782cc6f6e2a0fa5449656988f1671a6f4d972566dc370f5da4a455ae5d"
//
//		hash := hex.EncodeToString(byteHash)
//		if hash != expectedHash {
//			t.Error("hash is not expected", hash, expectedHash)
//		}
//	}
//
//	func TestTreeOddContent(t *testing.T) {
//		opts := MerkleTreeOpts{
//			HashFunc: DefaultHashFunc,
//		}
//
//		contents := make([]Content, 0, 3)
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//
//		tree, err := NewMerkleTree(opts, contents)
//		if err != nil {
//			t.Error("construct err", err)
//		}
//
//		byteHash, err := tree.CalculateHash()
//		if err != nil {
//			t.Error(err)
//		}
//
//		expectedHash := "6650e826af5b7813c2172d28f93b8e54ed0cc91f4a5729681433974fa76d0333"
//
//		hash := hex.EncodeToString(byteHash)
//		if hash != expectedHash {
//			t.Error("hash is not expected", hash, expectedHash)
//		}
//	}
//
//	func TestTreeManyContent(t *testing.T) {
//		opts := MerkleTreeOpts{
//			HashFunc: DefaultHashFunc,
//		}
//
//		contents := make([]Content, 0, 18)
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("123"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("456"), opts.HashFunc))
//		contents = append(contents, NewDefaultContent([]byte("789"), opts.HashFunc))
//
//		tree, err := NewMerkleTree(opts, contents)
//		if err != nil {
//			t.Error("construct err", err)
//		}
//
//		byteHash, err := tree.CalculateHash()
//		if err != nil {
//			t.Error(err)
//		}
//
//		expectedHash := "74449a8744bf18b90815fb5cc89fef53feffb3449108f2158b6387e73511894c"
//
//		hash := hex.EncodeToString(byteHash)
//		if hash != expectedHash {
//			t.Error("hash is not expected", hash, expectedHash)
//		}
//	}
//
// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//
//	func RandStringBytes(n int) []byte {
//		b := make([]byte, n)
//		for i := range b {
//			b[i] = letterBytes[rand.Intn(len(letterBytes))]
//		}
//		return b
//	}
//
//	func TestTreeRandomString(t *testing.T) {
//		opts := MerkleTreeOpts{
//			HashFunc: DefaultHashFunc,
//		}
//
//		contents := make([]Content, 0, 1_000_000)
//		for i := 0; i < 1_000_000; i++ {
//			contents = append(contents, NewDefaultContent(RandStringBytes(40), DefaultHashFunc))
//		}
//
//		tree, err := NewMerkleTree(opts, contents)
//		if err != nil {
//			t.Error("construct err", err)
//		}
//
//		_, err = tree.CalculateHash()
//		if err != nil {
//			t.Error(err)
//		}
//
// }

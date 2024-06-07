package crypto

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	var pub [130]byte
	var pr [64]byte
	NewPair(&pr, &pub)
	// fmt.Println(string(pub[:]))
	// fmt.Println(string(pr[:]))
}

func BenchmarkGenerate(b *testing.B) {
	var pub [130]byte
	var pr [64]byte

	for i := 0; i < b.N; i++ {
		NewPair(&pr, &pub)
	}
}

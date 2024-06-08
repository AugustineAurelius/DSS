package ecdh

import (
	"bytes"
	"testing"
)

func BenchmarkGenerate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

func TestHandshake(t *testing.T) {

	key1 := New()
	key2 := New()

	sec1, _ := key1.ECDH(key2.PublicKey())

	sec2, _ := key2.ECDH(key1.PublicKey())

	if !bytes.Equal(sec1, sec2) {
		t.Error("secrets in not equal", sec1, sec2)
	}

}

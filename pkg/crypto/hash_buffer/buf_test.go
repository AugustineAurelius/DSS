package hashbuffer

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestSHA512(t *testing.T) {
	str := strings.Repeat("12311111", 10000)
	msg := []byte(str)
	buf := New()

	bufM := buf.Hash(msg)

	shaM := sha512.Sum512(msg)

	fmt.Println(hex.EncodeToString(bufM[:]), hex.EncodeToString(shaM[:]))
}

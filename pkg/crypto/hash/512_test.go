package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestSHA512(t *testing.T) {
	str := strings.Repeat("12311111", 10000)
	msg := []byte(str)

	bufM := Hash512(msg)

	shaM := sha512.Sum512(msg)

	fmt.Println(hex.EncodeToString(bufM[:]), hex.EncodeToString(shaM[:]))
}

func TestSHA512ConcurentNative(t *testing.T) {
	str := strings.Repeat("12311111", 100)
	msg := []byte(str)

	wg := sync.WaitGroup{}
	wg.Add(1_000_000)

	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			sha512.Sum512(msg)
		}()
	}

	wg.Wait()
}

func TestSHA512Concurent(t *testing.T) {
	str := strings.Repeat("12311111", 100)
	msg := []byte(str)

	wg := sync.WaitGroup{}
	wg.Add(1_000_000)

	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			Hash512(msg)
		}()
	}

	wg.Wait()

}

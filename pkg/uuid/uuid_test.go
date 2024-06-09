package uuid

import (
	"testing"
)

// one alloc
func BenchmarkNew(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_ = New()
	}

}

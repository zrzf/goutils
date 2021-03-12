package strings

import (
	"github.com/zrzf/goutils/semaphore"
	"testing"
)

func BenchmarkRemoveBetweenRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveBetweenRune("<abc>abc<abc>", '<', '>', "")
	}
}

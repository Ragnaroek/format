package benchmarks

import (
	"fmt"
	"testing"
)

func BenchmarkFmtSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%s", "foo")
	}
}

func BenchmarkFmtLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%d%s%v%d%d%d%s", 666, "foo", "debug", 6, 6, 6, "test string, a little longer")
	}
}

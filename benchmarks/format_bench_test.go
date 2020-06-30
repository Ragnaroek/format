package benchmarks

import (
	"testing"

	"github.com/ragnaroek/format/pkg"
)

func BenchmarkFormatSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ft.Sformat("~a", "foo")
	}
}

func BenchmarkFormatLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ft.Sformat("~d~a~a~d~d~d~a", 666, "foo", "debug", 6, 6, 6, "test string, a little longer")
	}
}

func BenchmarkFormatFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ft.Sformat("~5,2,,,'0f", 1.111)
	}
}

package main

// $ go test -bench=.
// BenchmarkPopCount-8             2000000000               0.30 ns/op
// BenchmarkPopCount25-8           50000000                24.5 ns/op
// PASS
// ok      github.com/koshigoe/the-go-programming-language/chapter2/popcount/ex_2_5        1.902s

import (
	"testing"
	"math/rand"
	"github.com/koshigoe/the-go-programming-language/chapter2/popcount"
)

var arg = rand.Uint64()

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(arg)
	}
}

func BenchmarkPopCount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountEx25(arg)
	}
}

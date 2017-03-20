package main

// $ go test -bench=.
// BenchmarkPopCount-8             2000000000               0.30 ns/op
// BenchmarkPopCount23-8           100000000               16.8 ns/op
// PASS
// ok      github.com/koshigoe/the-go-programming-language/chapter2/popcount/ex_2_3        2.337s

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

func BenchmarkPopCount23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountEx23(arg)
	}
}

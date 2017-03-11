// 練習問題 1.3 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、実行時間の性能の差を計測しなさい。

package main

import "testing"

func makeArgs() []string {
	var args []string
	for i := 0; i < 10000; i++ {
		args = append(args, " ")
	}
	return args
}

var args = makeArgs()

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(args)
	}
}

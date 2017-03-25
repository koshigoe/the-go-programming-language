package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(basename1(os.Args[1]))
}

// basename1 はディレクトリ要素と.接尾辞を取り除きます。
// e.g. a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename1(s string) string {
	// 最後の '/' とその前の全てを破棄する。
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 最後の '.' より前の全てを保持する。
	for i:= len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

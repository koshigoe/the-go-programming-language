package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(basename1(os.Args[1]))
	fmt.Println(basename2(os.Args[1]))
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

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // "/" が見つからなければ -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

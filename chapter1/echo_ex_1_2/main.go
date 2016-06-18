// 練習問題 1.2 echo プログラムを修正して、個々の引数のインデックスと値の組を1行ごとに表示しなさい。
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

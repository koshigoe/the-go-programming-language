package main

import (
	"fmt"
)

// squares は呼び出されるごとに次の平方数を返す関数を返します。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
}

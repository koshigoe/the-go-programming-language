// ck は、その数値引数を摂氏と絶対温度へ変換します。
package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/koshigoe/the-go-programming-language/chapter2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToK(c), k, tempconv.KToC(k))
	}
}

// noempty はスライス内アルゴリズムの例です。
package main

import "fmt"

// noempty は文字列ではない文字列を保持するスライスを返します。
// 規程配列は呼び出し中に修正されます。
func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// append を使って書く例
func noempty2(strings []string) []string {
	out := strings[:0] // 元の長さ 0 のスライス
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty2(data2))
	fmt.Printf("%q\n", data2)
}

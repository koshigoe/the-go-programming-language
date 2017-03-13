// tempconv パッケージは摂氏(Celsius)と華氏(Fahrenheit)の温度計算を行います。
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%g℃", c) }

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)       // "100" ℃
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC)) // "180" ℉
	// fmt.Printf("%g\n", boilingF - FreezingC)    // コンパイルエラー

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	// fmt.Println(c == f)       // コンパイルエラー
	fmt.Println(c == Celsius(f)) // "true"

	c2 := FToC(212.0)
	fmt.Println(c2.String()) // "100℃"
	fmt.Printf("%v\n", c2)   // "100℃"; String を明示的に呼び出す必要はない
	fmt.Printf("%s\n", c2)   // "100℃"
	fmt.Println(c2)          // "100℃"
	fmt.Printf("%g\n", c2)   // "100"; String を呼び出さない
	fmt.Println(float64(c2)) // "100"; String を呼び出さない
}

package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("%#v\n", w)
	// 出力:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42
	fmt.Printf("%#v\n", w)
	// 出力:
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:20}
}

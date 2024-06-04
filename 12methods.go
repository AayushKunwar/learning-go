package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type Foo int32

func (x Foo) Abs() int32 {
	return int32(x + 1)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{Y: 4, X: 3}
	fmt.Println(v.Abs())
	w := Foo(32)
	fmt.Println(w.Abs())
}

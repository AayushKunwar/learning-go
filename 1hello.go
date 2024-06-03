package main

import (
	"fmt"
	"math/big"
)

var c, python, java bool

func main() {
	const World = "foobar"
	fmt.Println(World)
	newBigInt := big.NewInt(int64(69))
	fmt.Println(newBigInt)
	// fmt.Println("Hello World!")
	// fmt.Println(math.Pi)
	// fmt.Println(add(2,3))
	// var i, j int = 2, 3
	// foo := 69
	// fmt.Println(foo)
	// var c, python, java = true, false, "no!"
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// fmt.Println(split(18))
	// fmt.Println(i, j)
	// fmt.Println(c, python, java)
	var i int = 42
	f := 69.420
	u := uint(f)
	fmt.Println(i, f, u)
}

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
package main

import "fmt"

func main() {
	// q := [10]int{2, 3, 5, 7, 8, 9}
	// fmt.Println(q)
	// fmt.Println(len(q), cap(q))
	// sli := q[0:10]
	// fmt.Println(sli)
	a := make([]int, 5, 8)
	a[0] = 69
	printSlice("a", a)
	a = a[:6]
	printSlice("b", a)
}

func printSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
package main

import "fmt"

func main() {
	var s []int
	s = append(s, 345)
	fmt.Println(s)
	foo := [2]int{1, 2}
	s = append(s, foo[0])
	fmt.Println(s)
}

package main

import (
	"fmt"
)

func foo(bar int) int {
	bar++
	return bar
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if foo := sum; foo%2 != 0 {
			fmt.Println(foo)
		}
	}
	fmt.Println(sum)
}

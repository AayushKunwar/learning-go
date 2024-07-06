package main

import "fmt"

// fibonacci using function closure
func fibonacci() func() int {
	var prev, curr = 1, 0
	return func() int {
		next := prev + curr
		prev = curr
		curr = next
		return prev
	}
}

func main() {
	const foo = 69
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

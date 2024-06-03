package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := 0; i < 10; i++ {
		pow[i] = 1 << i
	}
	for _, value := range pow {
		fmt.Println(value)
	}
}

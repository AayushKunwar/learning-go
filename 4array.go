package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	var foo = [4]int{1,2,3,4}
	fmt.Println(foo)

	primes := [6]int{2,3,5,6,77,12}
	var s []int = primes[1:4]
	fmt.Println(s)
	s[0] = 0
	fmt.Println(primes)
}
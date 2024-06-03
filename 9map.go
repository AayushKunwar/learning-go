package main

import "fmt"

var m = map[string]uint8{
	"foo": 12,
	"bar": 21,
}

func main() {
	fmt.Println(m["bar"])
}

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for _, v := range words {
		// result[v] = len(v)
		result[v]++
	}
	return result
}

func main() {
	fmt.Println(WordCount("foo bar foo baz"))
}

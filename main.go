package main

import (
	"fmt"
)

func main() {
	fmt.Println(IntMin(10, 20))
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sum(a, b int) int {
	return a + b
}

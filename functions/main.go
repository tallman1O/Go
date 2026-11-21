package main

import (
	"fmt"
	"math"
)

// Both are valid ways to declare a function
// func add(a, b int) int {}
// func add(int a, int b) int {}

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}
func main() {
	// fmt.Println(math.pi) 
	// // first letter of the variable name is capitalised because it is exported (visible outside the package)

	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(swap("hello", "world"))
}
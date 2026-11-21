package main

import "fmt"

func main() {
	// 1. Inline: declare and call immediately
	func(msg string) {
		fmt.Println("inline:", msg)
	}("hello")

	// 2. Assign to a variable, then call later
	greet := func(name string) string {
		return "hi, " + name
	}
	fmt.Println(greet("Go"))

	// 3. Closure: returned func keeps its own sum
	add := makeAdder()
	fmt.Println(add(1)) // 1
	fmt.Println(add(2)) // 3
	fmt.Println(add(5)) // 8

	other := makeAdder() // independent closure
	fmt.Println(other(10)) // 10
}

// makeAdder returns an anonymous function that closes over sum.
func makeAdder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

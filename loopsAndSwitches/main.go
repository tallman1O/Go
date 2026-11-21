package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 10
	// switch statement
	switch j {
	case 0:
		fmt.Println("i is 0")
	case 1:
		fmt.Println("i is 1")
	default:
		fmt.Println("i is not 0 or 1")
	}

	// switch with multiple cases
	switch j {
	case 0, 1, 2:
		fmt.Println("i is 0, 1, or 2")
	default:
		fmt.Println("i is not 0, 1, or 2")
	}
}
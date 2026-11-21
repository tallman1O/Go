package main

import "fmt"

// array is a fixed-size sequence of elements of the same type
// cannot be resized after declaration
func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [4]int{2: 2, 3: 3}
	fmt.Println(arr3)

	arr4 := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr4)
}
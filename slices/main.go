package main 

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	s = primes[:4]
	fmt.Println(s)

	s = primes[2:]
	fmt.Println(s)

	// slice literal
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]

	// make method to create a slice
	// make(type, length, capacity)
	a := make([]int, 5) // length and capacity are 5
	fmt.Println(a) // [0 0 0 0 0]

	b := make([]int, 0, 5) // length is 0 and capacity is 5
	fmt.Println(b) // []

	c := b[:2] // length is 2 and capacity is 5
	fmt.Println(c) // [0 0]

	d := c[2:5] // length is 3 and capacity is 3
	fmt.Println(d) // [0 0 0]

	// append function to add elements to a slice
	a = append(a, 1)
	fmt.Println(a) // [0 0 0 0 0 1]

	b = append(b, 2)
	fmt.Println(b) // [0 0 2]

	c = append(c, 3)
	fmt.Println(c) // [2 0 3]

	d = append(d, 4, 5, 6)
	fmt.Println(d) // [3 0 0 4 5 6]
	
}
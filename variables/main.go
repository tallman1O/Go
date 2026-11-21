package main

import "fmt"

func main() {
	//initialisation of variables
	var a int
	fmt.Println(a)

	// declare and initialise variables
	var b int = 10

	//or 
	c := 20 // type inference
	fmt.Println(b)
	fmt.Println(c)

	const d = 30
	// d = 20 // error: cannot assign to d (const int)
	fmt.Println(d)

	// data types 
	var i int
	var f float64
	var bo bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, bo, s)
	
}
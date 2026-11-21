package main

import "fmt"

//defer is used to execute a function after the current function returns
func main() {
	defer fmt.Println("world") // defer is used to execute a function after the current function returns
	fmt.Println("hello")
}
 
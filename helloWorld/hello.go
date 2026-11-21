package helloworld // helloworld is the name of the package (same basename as the file) that contains the hello function

import "fmt"

func Hello() { // Hello is a function that prints the string "Hello, World!" to the console
	fmt.Println("Hello, World!")
}
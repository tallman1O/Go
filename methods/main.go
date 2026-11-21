package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// A method looks like a function, but has a receiver before the name.
// Here (p Person) means: call this on a Person value.
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// Pointer receiver (*Person) can change the original struct.
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p := Person{Name: "Ada", Age: 30}

	// call like p.Method() — Go passes p as the receiver
	fmt.Println(p.Greet()) // Hi, I'm Ada
	fmt.Println("age:", p.Age)

	p.HaveBirthday() // changes Age because receiver is *Person
	fmt.Println("age after birthday:", p.Age)
}

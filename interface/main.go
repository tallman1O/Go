package main

import "fmt"

// An interface lists methods — any type with those methods "implements" it.
// No "implements" keyword; Go checks this automatically.
type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

// Person has Greet() → it implements Greeter
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// Dog has Greet() → it also implements Greeter
func (d Dog) Greet() string {
	return "Woof! I'm " + d.Name
}

// This function accepts ANY Greeter — Person, Dog, or anything else with Greet()
func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	person := Person{Name: "Ada"}
	dog := Dog{Name: "Rex"}

	sayHello(person) // Hi, I'm Ada
	sayHello(dog)    // Woof! I'm Rex

	// You can also store different types in one Greeter variable
	var g Greeter = person
	fmt.Println(g.Greet())
	g = dog
	fmt.Println(g.Greet())
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Vertex struct {
	X int
	Y int
}

func main() {
	person := Person{Name: "John", Age: 30}
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	person.Name = "Jane"
	person.Age = 25
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
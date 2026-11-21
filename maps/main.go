package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// lookup time is O(1)
func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat: 40.68433,
		Long: -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	v, ok := m["Bell Labs"]
	fmt.Println(v, ok) // ok is true if the key is in the map, false otherwise

	//literal map
	m = map[string]Vertex{
		"Google": {
			Lat: 1,
			Long: -2,
		},
	}
	fmt.Println(m)
	fmt.Println(m["Google"])

	//never use var (m map[string]Vertex) to declare a map, use make(map[string]Vertex) instead of liternal syntax - m = map[string]Vertex

}
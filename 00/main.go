package main // main is the name of the package that contains the main function

import (
	"fmt" // fmt is a package that contains the fmt.Println function
	"math/rand" // math is a package that contains the rand.Intn function
)

func main() { // main is the main function that is executed when the program is run

	fmt.Println("My Favourite Number is", rand.Intn(10)) // Intn is a capitalized to access the function from the math/rand package
}
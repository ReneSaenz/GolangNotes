package main

import "fmt"

// Define the structure
type rect struct {
	width, height int
}

// Define the methods of the structure
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return ((2 * r.width) + (2 * r.height))
}

func main() {

	// Create a structure based on the type
	r := rect{width: 10, height: 5}

	// Call the 'area' function of the structure
	fmt.Println("area: ", r.area())

	// Call the 'perim' function of the structure
	fmt.Println("Perimeter: ", r.perim())

	// Declare a pointer to the structure
	rptr := &r

	// You can use the dot notation with struct pointers.
	// Pointers are derefenced automatically
	fmt.Println("area: ", rptr.area())
	fmt.Println("Perimeter: ", rptr.perim())

}

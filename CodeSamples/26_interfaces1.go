package main

import "fmt"
import "math"
import "reflect"

// Declare interface
type geometry interface {
	area() float64
	perim() float64
}

/*

To implement interface, we need to implement all methods in the interface.
Here we implement geometry interface on structure rect

*/

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return ((2 * r.width) + (2 * r.height))
}

//--------------------------------------

type circle struct {
	radius float64
}

// Here we implement geometry interface on structure rect
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g is of type:", reflect.TypeOf(g))
	fmt.Println("Struct values:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	/*
	   The circle and rect struct types both implement the
	   geometry interface so we can use instances of these
	   structs as arguments to measure

	*/

	measure(r)
	measure(c)

}

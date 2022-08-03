package main

import "fmt"

/*
Other languages have classes.
Go uses struct to represent the same contruct.
local memory/execution stack
*/

// Struct definition: type <name> struct
// One field per line
// Public fields start with a cap letter
type Person struct {
	name     string
	Location string
}

func main() {
	// instantiating an object based on the struct
	// this variable will be on the local memory/execution stack
	p1 := Person{}

	// assign a value to the fields
	p1.name = "name1"

	println(p1.name)

	// You can instantiate in the same line, following the order to the fields
	p2 := Person{"name2", "US"}
	println(p2.name)

	// =================================
	// Create larger objects in the heap
	p := new(Person)
	p.name = "Person object in the heap"
  p.Location = "Person object location"

  fmt.Println("Person name: ", p.name)
  fmt.Println("Person location: ", p.Location)
	fmt.Println(p)

	// =================================
	// Create a reference type. Put ampersand in front of the type
	pref := &Person{"Person object in the heap", "US"}
	fmt.Println(pref)

}

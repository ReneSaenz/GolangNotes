package main

import "fmt"

/*
In other languages, a special kind of method that lives inside a class,
is called a constructor.
In Go, we use a function
*/

type myStruct struct {
	myMap map[string]string
}

// constructor function
// naming convention is keyword 'new' followed by the name of the struct
// It is better to return a pointer to the new struct
// NOTE: constructor functions take parameters to initialize the objects (structs)
func newMyStruct() *myStruct {
	//create the struct
	s := myStruct{}
	//create and initialize the map
	s.myMap = map[string]string{}
	// return the address of the new struct, as indicated by the function signature
	return &s
}

func main() {
	// Call the construtor function
	foo := newMyStruct()
	foo.myMap["bar"] = "baz"

	fmt.Println(foo)
}

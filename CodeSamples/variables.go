package main


import(
"fmt"
"os"
"reflect"
)

/*

package-level variables => global variables

Variables names must start with underscore "_" or a letter
Variables are going to be initialized with a zero for numeric
and an empty string for string variables

*/

/*
var(
name string      
course string    
module float64   
)
*/

var(
name, course, module = "rene", "Go fundamentals", 3.2
)

func main() {

/*

- Short assignment
- using := is declaring and initializing variables it only works inside functions
Thus, declare function variables using :=
- variables declared inside functions, must be used
- Go can use type inference. You do not need to specify the variable type

*/

    name := "rene";
    name = os.Getenv("USER");
    course := "Docker Deep Dive";
    module := 3.2;
    addr := &module;

	fmt.Println("var name is set to :", name , "and is of type:", reflect.TypeOf(name));
	fmt.Println("var course is set to :", course , "and is of type:", reflect.TypeOf(course));
	fmt.Println("var module is set to:", module, "and is of type:", reflect.TypeOf(module));
	fmt.Println("var module address is", &addr)
	fmt.Println("what is inside mem address", addr, "is value:", *addr)


}
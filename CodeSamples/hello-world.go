package main

/*

If we want to run the program as a standalone, the 

package main

need to be the first line

*/


import(
"fmt" 
"runtime"
)

/*

Program entry point. Only for executables
Takes no arguments
No return values

*/

func main() {

    fmt.Println("Hello W0rld from", runtime.GOOS)

}


/*

- Go is case sensitive
- Package fmt is part of standard library
- If we want our functions to be exposed or public, we need to capitalize the first letter.

So 

fmt.println != fmt.Println


*/
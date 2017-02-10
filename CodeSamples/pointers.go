package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}


func main() {
    i := 1

    // Passed by value
    fmt.Println("initial:", i)
    zeroval(i)

    fmt.Println("zeroval:", i)

    // Passed by reference
    // The &i syntax gives the memory address of i, i.e. a pointer to i.
    zeroptr(&i)

    // Pointers can be printed too.    
    fmt.Println("zeroptr:", i)
    fmt.Println("pointer:", &i)

    //----------------------------

    a := 43
    fmt.Println("a - ", a)
    fmt.Println("a's memory address - ", &a)

    var p *int = &a
    fmt.Println("p int pointer - ", p)

    // The above code makes p a pointer to the memory address where an int is stored.
    // Thus, p is of type "int pointer"

    // invalid
    //var c int = &a
    //fmt.Println("c => ", c)

    // Deference the pointer, so we can get the value at that memory address
    fmt.Println("int pointer p deref => ", *p)

    var b *int = &a
    fmt.Println("Change the value of a via a pointer to 42")
    *b = 42 // Change the value at this address to 42
    fmt.Println("the new value of a => ", a)
}
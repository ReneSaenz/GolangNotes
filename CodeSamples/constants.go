package main

import "fmt"

/*
iota
Within a constant declaration, the predeclared identifier iota
represents successive untyped integer constants
 */


const (
	A = iota  // 0
	B = iota  // 1
	C = iota  // 2
)

// iota gets reseted
const (
	D = iota  // 0
	E         // 1
	F         // 2
)

const (
	_ = iota  // 0
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)



func main(){
	fmt.Printf("A => %d\n", A)
	fmt.Printf("B => %d\n", B)
	fmt.Printf("C => %d\n", C)
	fmt.Printf("D => %d\n", D)
	fmt.Printf("E => %d\n", E)
	fmt.Printf("F => %d\n", F)

	fmt.Println("\nbinary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "docker course"
	author := "rene"

	fmt.Println(converter(author, module))

	numTerms, sum := add2(1, 3, 5, 9)
	println("Added:", numTerms, " terms for a total of:", sum)

}

func converter(author string, module string) (ret1 string, ret2 string) {
	author = strings.ToUpper(author)
	module = strings.Title(module)

	return author, module
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

/*
 Named return values: improve readability of functions
 Noticed that we used/modified the names of the return values
 in the function.
 NOTE: the return statement does not need an expression.
*/
func add2(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}

	numTerms = len(terms)
	return
}

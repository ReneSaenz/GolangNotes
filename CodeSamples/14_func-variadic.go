package main

import (
	"fmt"
)

/*

Variadic functions are functions that allow an arbitrary set of parameters,
instead of a fixed number of parameters

*/

func main() {

	displayMessage("Hello", "from", "my", "little", "friend")

	min := minimum(13, 12, 17, 14, 12, 7, 5, 2)
	fmt.Println("The minimum value is:", min)

}

func displayMessage(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

func minimum(values ...int) int {

	min := values[0]

	for _, i := range values {
		if i < min {
			min = i
		}
	}

	return min
}

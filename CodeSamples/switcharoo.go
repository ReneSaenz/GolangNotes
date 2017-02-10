package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

switch <simple statement> ; <expression> {
	case <expression> : <code>
	case <expression> : <code>
	default : <code>
}


- switch type and case type, have to match
- the default case can go anywhere
- the switch stmt does not need a break stmt. Each case has an implicit break
- if we want the next case stmt to execute, use fallthrough keyword. Works in a case-by-case
- case stmt can have multiple comma-separated values
*/

func main() {

	switch "docker" {
	case "linux":
		fmt.Println("Linux")
	case "docker":
		fmt.Println("Docker")
		fallthrough
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Sorry, could not find a match")
	}

	fmt.Println("\n")

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We have an even num:", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We have an odd num:", tmpNum)
	}

}

func random() int {

	rand.Seed(time.Now().Unix())
	return rand.Intn(10)

}

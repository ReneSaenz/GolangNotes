package main

import (
	"fmt"
)

/*

boolean expressions for conditionals

==  Equals
!=  Not equals
<   Less than
<=  Less than or equal to
>   Greater than
>=  Greater than or equal
&&  AND
||  OR

*/

/*

if <boolean expression> {
	<code block>
} else if <boolean expression> {
	<code block>
} else {
	<code block>
}

*/

/*

Simple statements. Used to initialize variables used in the code block

if <simple statement>; <boolean expression> {
	<code block>
}

*/

func main() {

	val1 := 10
	val2 := 20

	if val1 < val2 {
		fmt.Println("Value1 is less than Value2")
	} else if val1 > val2 {
		fmt.Println("Value1 is greater than Value2")
	} else {
		fmt.Println("Value1 and Value2 are equal")
	}

	// declare variables in the if statement. Variable scope is within the if stmt.
	if v1, v2 := 30, 30; v1 < v2 {
		fmt.Println("Value1 is less than Value2")
	} else if v1 > v2 {
		fmt.Println("Value1 is greater than Value2")
	} else {
		fmt.Println("Value1 and Value2 are equal")
	}
}

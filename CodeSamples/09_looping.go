package main

import (
	"fmt"
	"time"
)

/*

Go only has the for key to perform loops

for <expression>


- Infinite loop
for{
	<code block>
}


- Conditional
for i < 10 {
	<code block>
}

for i := 0; i < 10; i++ {
	<code block>
}


- Range
for i := range courseList {

}


*/

func main() {

	// runCountDown()

	runRangeLoop()

}

func runCountDown() {
	for timer := 10; timer >= 0; timer-- {

		if timer == 0 {
			fmt.Println("ka-Boom!")
			//break
		} else {
			fmt.Println("Timer:", timer)
			time.Sleep(1 * time.Second)
		}

	}
}

func runRangeLoop() {
	courses := []string{"Docker Deep Dive",
		"Docker Clustering",
		"Docker and Kubernetes"}

	courseCompleted := []string{"Docker Deep Dive",
		"Go Fundamentals",
		"Puppet Fundamentals"}


    // Range return two values: index, data
	for _, i := range courses {
		// fmt.Println(i)
		for _, j := range courseCompleted {
			if i == j {
				fmt.Println("We found a match:", j)
			}
		}
	}

}

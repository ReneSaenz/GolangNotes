package main

import (
	"fmt"
	"reflect"
)

func main() {
	example1()
}

func example1() {
	//struct definition
	type courseMeta struct {
		author string  // struct field
		level  string  // struct field
		rating float64 // struct field
	}

	var DockerDeepDive courseMeta
	var OOP = new(courseMeta)

	Programming := courseMeta{
		author: "nene",
		level:  "intermediate",
		rating: 5,
	}

	fmt.Println(DockerDeepDive)

	fmt.Println(OOP)

	fmt.Println(Programming)

	DockerDeepDive.author = "Nigel Pulton"
	DockerDeepDive.rating = 5

	fmt.Println(DockerDeepDive)

	OOP.author = "viri"
	OOP.level = "beginner"
	OOP.rating = 5

	fmt.Println(OOP)

	fmt.Println("var DockerDeepDive is of type:", reflect.TypeOf(DockerDeepDive))
	fmt.Println("var Programming is of type:", reflect.TypeOf(Programming))
	fmt.Println("var OOP is of type:", reflect.TypeOf(OOP))

}

package main


import(
"fmt"
"reflect"
)



func main(){
	example1()
}


func example1(){

	type courseMeta struct {
		author string
		level string
		rating float64
	}

	var DockerDeepDive courseMeta
	var OOP = new(courseMeta)

	Programming := courseMeta {
		author: "nene",
		level: "intermediate",
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


    fmt.Println("var DockerDeepDive is of type:", reflect.TypeOf(DockerDeepDive));
    fmt.Println("var Programming is of type:", reflect.TypeOf(Programming));
    fmt.Println("var OOP is of type:", reflect.TypeOf(OOP));

}

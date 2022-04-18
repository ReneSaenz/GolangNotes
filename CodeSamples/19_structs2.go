package main

import "fmt"


type Person struct {
	Name string
}


type Saiyan struct {
	*Person
	Power int
}


func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}


func main() {

	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power: 9001,
    }

    goku.Introduce();

    // nene := Saiyan{
    // 	Person: 
    // 	Power: 1000
    // }

    // fmt.Println("Goku:", *goku.Person)
    // fmt.Println("Power:", goku.Power)


    // fmt.Println("\n")

    // fmt.Println("Nene Struct")\
    // fmt.Println("Name", )
}
	
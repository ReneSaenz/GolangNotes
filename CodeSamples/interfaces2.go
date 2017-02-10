package main


import "fmt"


/*

Define an Animal type that has a method named Speak

NOTE: Any type that defines this method is said to satisfy the Animal interface
There is no implements keyword. Whether or not a type satisfies an interface is
 determined automatically.

 */


type Animal interface{
	Speak() string
}

//---------------------------
type Dog struct{ }

func (d Dog) Speak() string{
	return "woof"
}


//---------------------------
type Cat struct{ }

func (c Cat) Speak() string{
	return "Meaow"
}


//---------------------------
type Llama struct { }

func (l Llama) Speak() string{
	return "???"
}


//---------------------------
type JavaProgrammer struct { }

func (j JavaProgrammer) Speak() string{
	return "Design Patterns"
}


func main(){
	makeSpeach()
}


func makeSpeach(){
	animals := []Animal{ Dog{}, Cat{}, Llama{}, JavaProgrammer{}}

	for _, animal:= range animals{
		fmt.Println(animal.Speak())
	}

}
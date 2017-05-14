package main

import "fmt"

type Person struct {
	name     string
	Location string
}

func NewPerson(name string) *Person {
	p := Person{
		name: name,
	}

	return &p
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	p := NewPerson("nene")
	fmt.Printf("Name: %s\n", p.Name())

	p.Location = "Texas"
	fmt.Printf("Location: %s\n", p.Location)

	p.SetName("rene")
	fmt.Printf("Name: %s\n", p.Name())
}

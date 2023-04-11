package main

import "fmt"

type HasName interface {
	GetNama() string
}

type Iperson struct {
	name string
}

func main() {

	iperson := Iperson{name: "adli"}
	sayHello(iperson)

}

func sayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetNama())
}

func (person Iperson) GetNama() string {
	return person.name
}

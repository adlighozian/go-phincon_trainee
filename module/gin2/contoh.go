package main

import "fmt"

func belajar() {

	c := new(int)

	// var b string = "test"
	
	var a *string 
	*a = "test"

	*c = 90

	// var numberA int = 4
	// var numberB *int = &numberA

	// *a = "bebas"

	// fmt.Println(a)
	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)

	*c = 10

	fmt.Println(c)
	fmt.Println(&c)
	fmt.Println(*c)

	// fmt.Println(&a)
}

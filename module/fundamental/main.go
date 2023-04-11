package main

import "fmt"

func main() {
	type coba string
	var belajar coba = "ini dia"
	// var nilai int = 6
	const text = `Aest`
	const a bool = true
	const (
		red        = iota
		orange int = 6
		green = iota
		yellow	
	)

	// nama = "adli"
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Printf("The value of Red is %v %T \n", nilai32, nilai32)
	fmt.Println(red)
	fmt.Println(orange)
	fmt.Println(yellow)
	fmt.Println(green)
	fmt.Println("===================")
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println("===================")
	fmt.Println(byte(text[0]))
	fmt.Println(a)
	fmt.Println(belajar)
	// fmt.Println(text)
	// fmt.Println(nama)
	// fmt.Println("hello world")
}

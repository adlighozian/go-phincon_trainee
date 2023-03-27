package main

import "fmt"

type Company struct {
	nama   string
	alamat string
	link   string
}

type Person struct {
	nama   string
	alamat string
	Company
}

func main() {
	fmt.Println("welcome to day 4")

	// func cli
	fmt.Println()
	fmt.Println("==========func cli==========")

	// var a int
	// var b int

	// fmt.Println("silahkan masukan nilai a")
	// fmt.Scanln(&a)
	// fmt.Println("silahkan masukan nilai b")
	// fmt.Scanln(&b)

	// fmt.Printf("nilai a = %d, nilai b = %d, hasilnya adalah = %d", a, b, hitung(a, b))

	// struct
	fmt.Println()
	fmt.Println("==========struct==========")

	type Customer struct {
		nama  string
		ktp   int32
		noHp  int32
		absen bool
	}

	var customer1 Customer
	customer1.nama = "nama saya bagas"
	customer1.ktp = 928329
	customer1.absen = true

	customer2 := Customer{"ini", 1, 7, true}

	fmt.Println("customer 1", customer1)
	fmt.Println("customer 2", customer2)

	// embed struct
	fmt.Println()
	fmt.Println("==========embed struct==========")

	var adli Person
	adli.nama = "adli"
	adli.Company.alamat = "jl. menjangan"

	var bagus Person
	bagus.nama = "bagus"

	reslut1 := bagus.nama

	fmt.Println(reslut1)

	// method struct
	fmt.Println()
	fmt.Println("==========method struct==========")

	adli.sayHello()
	// funcstruct_ = func (customer Person) sayHello()

}

// func struct method
func (person Person) sayHello() {
	println("hello, my name is ", person.nama)
}

// func cli
func hitung(a, b int) int {
	parA := a
	parB := b
	result := parA + parB

	return result

}

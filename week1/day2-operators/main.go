package main

import (
	"fmt"
)

func main() {
	// operator matematika
	a := 3
	b := 5
	tambah := a + b
	kurang := b - a

	fmt.Println("==========Operator matematika==========")
	fmt.Println(float32(b) / float32(a))
	fmt.Println("pertambahan", tambah)
	fmt.Println("pengurangan", kurang)

	// Assignment operator
	c := 15
	d := 20
	var e string = "coba"
	c++
	e += "test"

	fmt.Println()
	fmt.Println("==========Assignment operator==========")
	fmt.Println(c)
	fmt.Println(d)

	// Relational operator
	nama1 := "bagas"
	nama2 := "nandy"
	result := nama1 == nama2
	name3 := "A"
	name4 := "a"

	fmt.Println()
	fmt.Println("==========Relational operator==========")
	fmt.Println(result)
	fmt.Println(name3 > name4)

	// Logical operator
	var left bool = false
	var right bool = true
	var leftAndRight bool = left && right
	var rightAndLeft bool = left || right
	var leftReserve bool = !left

	fmt.Println()
	fmt.Println("==========Logical operator==========")
	fmt.Println(leftAndRight)
	fmt.Println(rightAndLeft)
	fmt.Println(leftReserve)

	// Pecabangan if
	var ifa string = ""
	name := "fian"
	if name == "adli" {
		ifa = "benar"
	} else if name == "fian" {
		ifa = "salah"
	} else {
		ifa = "siapa nih?"
	}

	fmt.Println()
	fmt.Println("==========Percabangan if==========")
	fmt.Println(ifa)

	// Temporary variable
	namet := "adli"
	if len := len(namet); len > 5 {
		fmt.Println()
		fmt.Println("==========Temporary variable==========")
		fmt.Println("nama terlalu panjang", len)
	} else {
		fmt.Println()
		fmt.Println("==========Temporary variable==========")
		fmt.Println("nama sudah benar", len)
	}

	//Switch expression
	names := "adli"
	result1 := ""

	switch names {
	case "adli":
		result1 = "Hello adli"
	}

	fmt.Println()
	fmt.Println("==========Switch expression==========")
	fmt.Println(result1)

	//for loop
	fmt.Println()
	fmt.Println("==========for loop==========")

	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for counters := 1; counters <= 10; counters++ {
		fmt.Println("(with statement) perulangan ke", counters)
	}

	//for range
	fmt.Println()
	fmt.Println("==========for range==========")

	datas := []int{1, 2, 3} // slice
	datas1 := [...]string{"ada", "ini", "itu"}

	for index, value := range datas {
		fmt.Println("Index", index, "=", value)
	}

	// Latihan 1
	fmt.Println()
	fmt.Println("==========Latihan 1==========")
	var nilai1 int = 70
	resultif := ""
	resultswitch := ""
	message := "Anda mendapatkan nilai"

	if nilai1 >= 90 {
		resultif = "A"
	} else if nilai1 >= 80 {
		resultif = "B"
	} else if nilai1 >= 70 {
		resultif = "C"
	} else if nilai1 >= 60 {
		resultif = "D"
	} else {
		resultif = "E"
	}

	fmt.Println(message, resultif)
	fmt.Println(resultswitch)

	//latihan 2
	fmt.Println()
	fmt.Println("==========latihan 2==========")

	price := 27
	qty := 2
	total := price * qty
	disc := true

	switch disc {
	case true:
		fmt.Println(float32(total) / 10)
	default:
		fmt.Println(total)
	}

	//latihan 3
	fmt.Println()
	fmt.Println("==========latihan 3==========")

	counts3 := 0
	jmlstr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	for _, value := range jmlstr {

		switch string(value) {
		case "A", "a", "I", "i", "U", "u", "E", "e", "O", "o":
			counts3++
		}
	}
	fmt.Println(counts3)

	//latihan 4
	fmt.Println()
	fmt.Println("==========latihan 4==========")

	var bil float64 = 30
	var wadah float64 = 1
	// var kotak int

	for bil >= 1 {
		wadah = wadah * bil
		bil--
	}
	fmt.Println(wadah)

	//latihan 5
	fmt.Println()
	fmt.Println("==========latihan 5==========")

	bilangan := "satu"

	switch bilangan {
	case "satu", "tiga", "lima", "tujuh", "sembilan":
		fmt.Println("ganjil")
	case "dua", "empat", "enam", "delapan", "sepuluh":
		fmt.Println("genap")
	default:
		fmt.Println("tidak ditemukan")
	}

	//  continue & break
	fmt.Println()
	fmt.Println("==========continue & break==========")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// trial
	fmt.Println()
	fmt.Println("==========trial==========")
	for index, value := range datas1 {
		fmt.Println(index, string(value))
	}

}

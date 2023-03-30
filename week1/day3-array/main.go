package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Array
	fmt.Println()
	fmt.Println("==========Array==========")

	var values = [3]int{
		12, 13, 25,
	}

	var arrstr = [3]string{
		"halo", "ini", "itu",
	}
	fmt.Println(values)
	fmt.Println(arrstr)
	fmt.Println("jumlah element", len(values))
	values[0] = 100
	fmt.Println(values[0])

	// Data slice
	fmt.Println()
	fmt.Println("==========Data slice==========")

	months := []string{
		"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november,", "desember",
	}

	sliceMake := []rune("test")
	// letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	slice1 := months[4:7]
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice2 := months[6:9]
	fmt.Println(slice2)
	slice1[0] = "test"
	fmt.Println(months)

	var arr1 [1]string
	arr1[0] = "bagas"
	arr1[0] = "rafly"

	var arr2 = [9]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}
	// arr2[0] = 8

	slicetest := arr2[1:4]
	fmt.Println("slice 1:4 ", slicetest)
	slicetest = append(slicetest, 11)

	slicetest2 := slicetest[2:3]
	fmt.Println("slice 2:3 ", slicetest2)

	slicetest2 = append(slicetest2, 21)

	// maket := make([]string, 2, 3)

	test := make([]int, 5, 6)
	copy(test, slicetest2)

	fmt.Println("array =", arr2, len(arr2), cap(arr2))
	fmt.Println("slice 1 =", slicetest, len(slicetest), cap(slicetest))
	fmt.Println("slice 2=", slicetest2, len(slicetest2), cap(slicetest2))
	// fmt.Println(s, len(s), cap(s))
	fmt.Println()
	fmt.Println(test)

	// latihan 1
	fmt.Println()
	fmt.Println("==========latihan 1==========")

	type hobi []string

	var person2 = make(map[string][]string)
	person2["adli"] = hobi{"berenang", "belajar", "bermain"}
	person2["bagas"] = hobi{"berenang", "belajar", "bermain"}
	// person2["hobi"] = hobi{"berenang", "belajar", "bermain"}

	// var wadah string
	// for i := 0; i < len(person2["adli"]); i++ {
	// 	wadah += person2["adli"][i] + " "
	// }

	for key, value := range person2 {
		var hobi string
		panjang := len(value)

		for i, v := range value {
			switch {
			case panjang == 1:
				hobi += fmt.Sprintf("%s.", v)
			case i == (panjang - 1):
				hobi += fmt.Sprintf("dan %s.", v)
			default:
				hobi += fmt.Sprintf("%s, ", v)
			}
		}
		fmt.Printf("nama saya %s mempunyai hobi %s\n", key, hobi)
	}

	// Map
	fmt.Println()
	fmt.Println("==========Map==========")

	mapBaru := make(map[string]string)

	var map2 = make(map[string]int)
	map2["nama"] = 235135

	// map2["nama"] = 0

	var person = map[string]string{
		"nama": "adli",
	}

	var map1 map[string]string = map[string]string{
		"nama": "adli",
		"umr":  "12",
	}

	// map2["id"] = 1

	mapBaru["id"] = "1"
	mapBaru["nama"] = "ini fikri"
	delete(mapBaru, "nama")

	fmt.Println(mapBaru)
	fmt.Println(map1["umr"])
	fmt.Println(person)
	fmt.Println(map2)

	// function
	fmt.Println()
	fmt.Println("==========function==========")

	belajar_func("adli", "ghozian")

	resultf := belajar_func2(1)
	resultf = (resultf - 2) * 23 * (-1)

	fmt.Println(resultf)

	// function multiple return
	fmt.Println()
	fmt.Println("==========function multiple return==========")

	firstname, username := func_multiple()
	fmt.Println(firstname, username)

	afirstname, _ := func_multiple()
	fmt.Println(afirstname)

	// function multiple return named
	fmt.Println()
	fmt.Println("==========function multiple return named==========")

	asad, message := func_multiplenamed("test")
	fmt.Println(asad, message)

	fmt.Println()
	fmt.Println("==========variadic function==========")

	vari := variadic(1, 2, 3)
	println(vari)

	varuslice := []int{2, 2, 2}
	fmt.Println(variadic(varuslice...))

	fmt.Println()
	fmt.Println("==========func as var==========")

	var funcvar func(int, int) float64 = pythagoras
	resultFuncVar := funcvar(2, 4)
	fmt.Println(resultFuncVar)

	fmt.Println()
	fmt.Println("==========func as param and annonmous fanc==========")

	// trial
	fmt.Println()
	fmt.Println("==========trial==========")

	pang := func_multiple_param

	hasilPenjumlahan := pang("d", "s", "a")

	fmt.Println(pythagoras(2, 3))
	fmt.Println(hasilPenjumlahan)
	fmt.Println(pang("ini", "punya", "siapa"))

	var wadah int
	sliceTest := []string{"test", "ad asdas", "ad adas"}
	for _, value := range sliceTest {

		if strings.Contains(strings.ToLower(value), "ad") {
			wadah++
		}

		fmt.Println(value, wadah)
	}

	// Latihan 2
	fmt.Println()
	fmt.Println("==========latihan 2==========")

	funclat := latihan
	fmt.Println(funclat("Aditya Ananta Putra",
		"ADNAN NUR JAILANI",
		"Afrila Zahra Prasetyo",
		"Aida Nirmala",
		"Amalia Rahma",
		"ANDINI DWI RIZKY PUTRI",
		"Angga Putra",
		"CAHAYA WIJAYA IJAM",
		"Dea Christina",
		"DEVITA NANDA OKTAVIA",
		"DEWI AYU LESTARI",
		"Dhevina Ananda Fitri",
		"DHEVY YANI RIZKI",
		"Dwi Mahani",
		"Eri Santosos",
		"Faiza Amalia Mahfudz",
		"FAUZIYATUNNISA",
		"Febriyanti Syabina",
		"Hafifah Nur Azmi Pratiwi",
		"Juliansyah Husien",
		"Kharisa Nur Aziza",
		"Lun Wina",
		"MAHREVA ROESTHA RAMADANIATY",
		"MUHAMMAD FADHILAH",
		"Muhammad Fatah Firdaus",
		"Muhammad Rafly Ahya",
		"Muhammad Rahmin Noor",
		"Muhammad Rivaldi",
		"Muhammad Rizky Rachmadhani",
		"Nadiah Nahdhah Islamiyah",
		"NOLA FEBRIANA SAPUTRI",
		"RAFADYAN REYHAN MARITZA WERAT",
		"Robby Satria",
		"Salshabilla Wahyu Anafi",
		"STEVI VIONI PAKULLA",
		"Winny Rahmah Nia"))

	fmt.Println()
	fmt.Println("==========func as param and annonmous fanc==========")

	// annonymous function
	filter_ := func(param string) bool {
		if strings.Contains(strings.ToLower(param), "anjing") {
			return true
		} else {
			return false
		}
	}

	phrase := "test anjing"
	printSpeach(phrase, filter_)

	fmt.Println(len(sliceMake))
}

// func as a param

type Filter func(string) bool

func printSpeach(speech string, filter Filter) {
	if filter(speech) {
		fmt.Println("mengandung kata-kata kasar")
	} else {
		fmt.Println("tidak mengandung kata-kata kasar")
	}
}

// func filteress(speech string) bool {
// 	if strings.Contains(strings.ToLower(speech), "anjing") {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// =================

// latihan
func latihan(param ...string) (int, []string, map[string]int) {

	var jumlahData int
	var jumlahAk int

	sliceMuhammad := make([]string, jumlahData)
	jumlaha := make(map[string]int)

	for _, nama := range param {
		jumlahData++
		if strings.Contains(strings.ToLower(nama), "muhammad") {
			sliceMuhammad = append(sliceMuhammad, nama)
		}
		jumlahAk = strings.Count(nama, "a")
		jumlaha[nama] = jumlahAk
	}

	return jumlahData, sliceMuhammad, jumlaha
}

func variadic(numbers ...int) int {
	jum := 0

	for _, val := range numbers {
		jum += val
	}
	return jum
}

func func_multiplenamed(param string) (coba string, message string) {
	coba, message = param, "jangan"
	return
}

func func_multiple_param(satu, dua, tiga string) string {
	return "hallo " + satu + dua + tiga
}

func func_multiple() (string, int) {
	return "ghoiz", 12
}

func belajar_func(namadepan string, namabelakang string) {
	fmt.Println("hello", namadepan, namabelakang)
}

func belajar_func2(umurb int) int {
	return umurb
}

func pythagoras(a int, b int) float64 {
	return math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
}

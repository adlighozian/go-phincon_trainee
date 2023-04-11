package main

import (
	"fmt"
)

func main() {
	var a any = "test"
	stringNama := a.(string)
	fmt.Println(stringNama)

	var num1 any = 1
	var num2 any = 2
	hasil := num1.(int) + num2.(int)
	fmt.Println(hasil)

	var sekolah any = []string{"adli"}

	for _, value := range sekolah.([]string) {
		fmt.Println(value)
	}

	// with switch
	var random any = map[string]string{"nama": "itu", "coba": "test"}
	switch result := random.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []string:
		fmt.Println("ini slice string")
	case []int:
		for _, v := range result {
			fmt.Println("nomor", v)
		}
	case map[string]string:

	default:
		println(result)
	}
}

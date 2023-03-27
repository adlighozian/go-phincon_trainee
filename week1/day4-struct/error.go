package main

import (
	"errors"
	"fmt"
)

func Pembagian(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("b tidak boleh 0")
	} else {
		return a / b, nil
	}
}

func main() {
	hasil, err := Pembagian(9, 0)
	if err == nil {
		fmt.Println("hasilnya", hasil)
	} else {
		fmt.Println(err.Error())
	}
}

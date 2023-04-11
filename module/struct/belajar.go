package main

import "fmt"

type Coba struct {
	Satu string
}

type Utama struct {
	Cobain int
	Coba
}

func main() {
	var kotak = Utama{}
	kotak.Coba.Satu = "ini embed"
	kotak.Cobain = 2

	fmt.Println(kotak.Coba.Satu)
	fmt.Println(kotak.Cobain)
}

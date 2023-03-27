package main

import "fmt"

type pointer struct {
	nama   string
	alamat string
	kucing string
}

func main() {

	alamat1 := pointer{
		nama:   "adli",
		alamat: "menjangan",
		kucing: "hamcing",
	}
	alamat2 := &alamat1

	fmt.Println(alamat1)
	fmt.Println(*alamat2)

	contoh := "luthfi"
	contoh1 := &contoh

	var contoh3 *string
	contoh3 = &contoh

	fmt.Println(contoh)
	fmt.Println(*contoh1)
	fmt.Println(contoh3)

	// ponter function
	nama1 := "dudi"
	fmt.Println("sebelum change name", nama1)
	changeName(&nama1)
	fmt.Println("sesduah change name", nama1)
}

func changeName(nama *string) {
	*nama = "bagas"
}

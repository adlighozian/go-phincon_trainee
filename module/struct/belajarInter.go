package main

import "fmt"

type dompet interface {
	isiDompet() string
}

type orang struct {
	uang string
	nama string
}

func (u orang) isiDompet() string {
	result := u.nama + " Mempunyai uang sebanyak " + u.uang
	return result
}

func lihatDompet(d dompet) {
	result := d.isiDompet()
	fmt.Println(result)
}

func main() {
	orang1 := orang{uang: "12", nama: "azril"}
	lihatDompet(orang1)
}

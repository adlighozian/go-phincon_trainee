package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup    //mengdeklarasikan waitgroup
var RWMutex sync.RWMutex //mendeklarasikan RWMutex

func f(v *int) { //membuat function f dengan parameter bertipedata pointer integer
	defer wg.Done()  // menjalankan wait group done di akhir
	RWMutex.Lock()   //Locking menggunakan RWMutex
	*v++             //Menambah v dengan +1
	RWMutex.Unlock() //Unlocking menggunakan RWMutex
}

func main() {
	var v int = 0 //mendeklarasikan v integer dengan nilai 0

	for i := 0; i < 10000; i++ { //perulangan for
		wg.Add(1) //untuk menandakan ada proses gorutine berjumlah 1
		go f(&v)  // menjalankan gorutine dengan function f yang membawa parameter v
	}

	wg.Wait()                  //menunggu sampai peroses selesai
	fmt.Println("Finished", v) //print hasil v
}

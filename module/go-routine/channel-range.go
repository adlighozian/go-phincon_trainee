package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	channel := make(chan string) // membuat channel baru

	go func() { //membuat annonymous function
		for i := 1; i <= 10; i++ { // perulangan
			channel <- "Perulangan ke " + strconv.Itoa(i) //untuk memasukkan value ke dalam channel
		}
		defer close(channel) // untuk menghentikan channel saat perulangan for sudah selesai
	}()

	for data := range channel {
		time.Sleep(1 * time.Second)
		fmt.Println(data, "range")
	}

	test := <-channel // kosong karena data dalam channel sudah habis

	fmt.Println("DONE")
	fmt.Println(test)
}

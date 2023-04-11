package main

import (
	"fmt"
)

func main() {

	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				x++
			}
		}()
	}

	fmt.Println("Jumlah counter", x)
}

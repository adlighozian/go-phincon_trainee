package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	var counter int
	
	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1", data)
		counter++
	case data := <-channel2:
		fmt.Println("Data dari channel 2", data)
		counter++
	default:
		fmt.Println("Menunggu Data")
	}

}

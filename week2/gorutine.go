package main

import (
	"fmt"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func message(number int) {

	fmt.Println("Display: ", number)
}

func main() {

	for i := 0; i < 10000; i++ {
		go message(i)
	}

	time.Sleep(5 * time.Second)
}

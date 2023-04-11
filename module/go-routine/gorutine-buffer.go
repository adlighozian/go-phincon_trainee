package main

import "fmt"

func main() {
	channel := make(chan string, 5)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

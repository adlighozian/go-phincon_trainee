package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	defer close(messages)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data

	}

	for i := 1; i <= 3; i++ {
		switch {
		case i == 1:
			go sayHelloTo("john wick")
		case i == 2:
			go sayHelloTo("ethan hunt")
		case i == 3:
			go sayHelloTo("jason bourne")
		}

	}

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

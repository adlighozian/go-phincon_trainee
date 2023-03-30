package main

import (
	"context"
	"fmt"
)

func main() {
	value := map[string]string{
		"1": "one",
		"2": "two",
	}
	background := context.Background()
	child := context.WithValue(background, "value", value)
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(child)
	fmt.Println(child.Value("value"))
	fmt.Println(todo)
	
}

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"nama": "adli",
		}
	}
}

func main() {
	data := NewMap("test")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}

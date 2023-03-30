package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// membuat randomizer dengan unix
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	// membuat slice untuk menampung letters dengan array of char
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// membuat slice dengan panjang dari parameter length
	b := make([]rune, 10)

	// di sini melakukan perulangan dari slice b dengan for range, setiap index b dimasukkan dengan nilai array of char dari letters yang sudah di random dengan randomizer
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
		fmt.Println(string(b))
		time.Sleep(1 * time.Second)
	}

	var a string = string(b)

	fmt.Println(a)
}

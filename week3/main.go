package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	// membuat slice untuk menampung letters dengan array of char
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// membuat slice dengan panjang dari parameter length dengan tipedara alias rune
	b := make([]rune, 7)

	// melakukan perulangan dari slice b, dan setiap index slice b di isi oleh value slice letters yang sudah di randomize
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	fmt.Println(string(b))
}

package helper

import (
	"math/rand"
	"time"
)

type RandomInterface interface {
	Randomizer() string
}

type random struct {
}

func NewRandom() *random {
	return &random{}
}

func (r *random) Randomizer() string {
	time.Sleep(1 * time.Second)
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 7)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	rand := string(b)
	return rand
}

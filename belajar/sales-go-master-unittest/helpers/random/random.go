package random

import (
	"math/rand"
	"strconv"
	"time"
)

type random struct {
}

func NewRandom() *random {
	return &random{}
}

type RandomInterface interface {
	RandomString(length int) (int, error)
}

func (r *random) RandomString(length int) (int, error) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	integer, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}

	return integer, nil
}
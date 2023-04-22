package random

import (
	"math/rand"
	"strconv"
)

func RandomString(length int) (int, error) {
	letters := []rune("1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	integer, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}

	return integer, nil
}
package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

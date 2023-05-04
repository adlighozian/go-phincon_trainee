package helper

import (
	"errors"

	"github.com/joho/godotenv"
)

func GetEnv(key string, callback string) (string, error) {
	err := godotenv.Load("C:\\Users\\user\\Documents\\GitHub\\go-phincon_trainee\\tugas\\contact-go-json\\.env")
	if err != nil {
		return callback, errors.New("error")
	}
	return key, nil
}

package middleware

import "log"

func FailError(err error, msg string) {
	if err != nil {
		log.Panicf("%s : %s", err, msg)
	}
}

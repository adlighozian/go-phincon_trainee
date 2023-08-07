package main

import (
	"log"
	"strings"
)

func main() {

	val, msg := check("ada ada aja")

	log.Println(val, msg)
}

func check(value string) (bool, string) {
	res := len(TrimLTSpace(value))
	
	return res == 0, TrimLTSpace(value)
}

func TrimLTSpace(value string) string {
	return strings.ReplaceAll(value, " ", "")
}

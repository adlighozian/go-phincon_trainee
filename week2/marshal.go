package main

import (
	"encoding/json"
	"fmt"
)

func encodeJSON() {
	name, _ := json.Marshal("Umar Bawazir")
	fmt.Println("JSON dari string", string(name))
	age, _ := json.Marshal(22)
	fmt.Println("JSON dari int", string(age))
	married, _ := json.Marshal(true)
	fmt.Println("JSON dari bool", string(married))
	hobbies, _ := json.Marshal([]string{"Gaming", "Ngodin", "Mancing"})
	fmt.Println("JSON dari slice", string(hobbies))
}

func main() {
	// Marshal

	type profile struct {
		nama   string
		umur   int
		alamat string
	}

	mars := profile{"adli", 12, "test"}
	b, _ := json.Marshal(mars)
	fmt.Println(b)
	fmt.Println()

	encodeJSON()
}

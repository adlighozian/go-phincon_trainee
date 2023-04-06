package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreeen() {
	var enter string
	fmt.Println("\n\nWant to choose another menu ? click enter !")
	fmt.Scanln(&enter)
	
	clearScreen := exec.Command("cmd", "/c", "cls")
	clearScreen.Stdout = os.Stdout
	clearScreen.Run()
}
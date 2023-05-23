package clearscreen

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreeen() {
	var enter string
	fmt.Println("\nPilih menu lain tekan enter")
	fmt.Scanln(&enter)

	clearScreen := exec.Command("cmd", "/c", "cls")
	clearScreen.Stdout = os.Stdout
	clearScreen.Run()
}
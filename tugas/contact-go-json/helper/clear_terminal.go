package helper

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

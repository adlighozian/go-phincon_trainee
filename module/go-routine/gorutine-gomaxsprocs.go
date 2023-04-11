package main

import (
	"fmt"
	"runtime"
)

func main() {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("CPU", totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("CPU", totalGoroutine)
}

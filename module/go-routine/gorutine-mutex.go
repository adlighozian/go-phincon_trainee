package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	var x int
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Jumlah counter", x)
}

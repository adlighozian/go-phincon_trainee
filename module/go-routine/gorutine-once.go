package main

import (
	"fmt"
	"sync"
)

var counter int

func OnlyOnce() {
	counter++
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			once.Do(OnlyOnce)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

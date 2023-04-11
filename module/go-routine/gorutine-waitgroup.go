package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsyncronous(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hellow", i)
	time.Sleep(1 * time.Second)
}

func main() {
	group := &sync.WaitGroup{}
	var mutex sync.RWMutex
	for i := 0; i < 10; i++ {
		mutex.Lock()
		go RunAsyncronous(group, i)
		mutex.Unlock()
	}

	group.Wait()
	fmt.Println("Complete")
}

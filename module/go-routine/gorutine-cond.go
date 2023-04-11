package main

import (
	"fmt"
	"sync"
	"time"
)

// Create a new condition variable with a mutex.
var cond = sync.NewCond(&sync.Mutex{})

// Create a WaitGroup.
var group = &sync.WaitGroup{}

// WaitCondition waits for a condition variable to be signaled.
func WaitCondition(value int) {
	cond.L.Lock() // Lock the condition variable's mutex.
	cond.Wait()   // Wait for the condition variable to be signaled.
	fmt.Println("Done", value)
	cond.L.Unlock() // Unlock the condition variable's mutex.
	group.Done()    // Notify the WaitGroup that this goroutine has finished.
}

func main() {
	for i := 1; i <= 10; i++ {
		group.Add(1)        // Add a new goroutine to the WaitGroup.
		go WaitCondition(i) // Start a new goroutine.
	}

	// Signal the condition variable 10 times.
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait() // Wait for all goroutines to finish.
}

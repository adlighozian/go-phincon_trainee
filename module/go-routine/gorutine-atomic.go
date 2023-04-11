package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// Atomic merupakan package yang digunakan untuk menggunakan
// data primitive secara aman pada proses concurrent
//
// safe and efficient atomic operations on variables
// shared across goroutines.
func main() {
    var group sync.WaitGroup
    var counter int64 = 0

    // create 100 goroutines
    for i := 0; i < 100; i++ {
        group.Add(1)

        go func() {
            // perform 100 increments on counter
            for j := 0; j < 100; j++ {
                atomic.AddInt64(&counter, 1)
            }

            group.Done()
        }()
    }

    group.Wait()

    // print final value of counter
    fmt.Println("counter:", counter)

}
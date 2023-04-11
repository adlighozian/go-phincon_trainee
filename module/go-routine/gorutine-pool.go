package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	for i := 0; i < 1; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	time.Sleep(1 * time.Second)
}

package main

import (
	"cari-nama-goroutine/helper"
	"cari-nama-goroutine/model"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	helper.CallClear()
	var wg sync.WaitGroup
	const timeoutDuration = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)

	defer cancel()

	wg.Add(1)
	go func() {
		readNamesWorker(ctx, &wg, model.Names)
		wg.Done()
	}()

	wg.Wait()
}

func readNamesWorker(ctx context.Context, wg *sync.WaitGroup, names []string) {
	numberOfWorkers := 10
	wg.Add(numberOfWorkers)
	go func() {
		fmt.Println("start check name with worker", numberOfWorkers)
		for workerIndex := 1; workerIndex <= numberOfWorkers; workerIndex++ {
			var total int
			go func(workerIndex int) {
				defer wg.Done()
				start := time.Now()
			L:
				for _, name := range names {
					select {
					case <-ctx.Done():
						break L
					default:
						if strings.Contains(strings.ToLower(name), "wina") {
							total++
						}
					}
				}
				
				duration := time.Since(start)
				fmt.Println("worker", workerIndex, "found", total, "of Wina. Done in", duration.Seconds(), "seconds")

			}(workerIndex)
		}

	}()

	go func() {
		wg.Wait()
		fmt.Println("end check name")
	}()
}

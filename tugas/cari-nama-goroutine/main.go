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

var wg = new(sync.WaitGroup)

func main() {
	helper.CallClear()

	const timeoutDuration = 5 * time.Hour

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	// wg.Add(2)

	getNames := getName(ctx, model.Names)
	filterName(getNames, ctx)
	var input string
	fmt.Scanln(&input)
	// wg.Wait()
}

func getName(ctx context.Context, names []string) <-chan string {
	chanOut := make(chan string, 5)

	go func() {
		select {
		case <-ctx.Done():
			break
		default:
			for _, v := range names {
				chanOut <- v
				// time.Sleep(1 * time.Second)
			}
		}
		defer close(chanOut)
	}()

	return chanOut
}

func filterName(chanIn <-chan string, ctx context.Context) {
	numberOfWorkers := 5
	wg.Add(numberOfWorkers)

	go func() {
		fmt.Println("start check name with worker", numberOfWorkers)
		for workerIndex := 1; workerIndex <= numberOfWorkers; workerIndex++ {
			var total int
			// time.Sleep(1 * time.Second)
			go func(workerIndex int) {
				defer wg.Done()
				start := time.Now()
			L:
				for name := range chanIn {
					// fmt.Println(name)

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

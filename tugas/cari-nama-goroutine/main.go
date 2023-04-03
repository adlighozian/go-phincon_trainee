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

	wg.Add(1)

	go func() {
		start := time.Now()
		getNames := getName(ctx, model.Names)
		filterName(getNames, ctx, start)

		wg.Done()
	}()

	wg.Wait()
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
			}
		}
		defer close(chanOut)
	}()

	return chanOut
}

func filterName(chanIn <-chan string, ctx context.Context, start time.Time) {
	numberOfWorkers := 5
	wg.Add(numberOfWorkers)

	go func() {
		fmt.Println("start check name with worker", numberOfWorkers)
		for workerIndex := 1; workerIndex <= numberOfWorkers; workerIndex++ {
			var total int
			go func(workerIndex int) {
				defer wg.Done()
			L:
				for name := range chanIn {
					// fmt.Println(name)
					select {
					case <-ctx.Done():
						break L
					default:
						if strings.Contains(strings.ToLower(name), "andreasti") {
							total++
						}
					}
				}

				duration := time.Since(start)
				fmt.Println("worker", workerIndex, "found", total, "duration", duration.Microseconds(), "seconds")
			}(workerIndex)
		}
		wg.Wait()
		fmt.Println("end check name")
	}()

}

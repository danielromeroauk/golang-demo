package main

import (
	"danielromeroauk/golang-demo/pkg/exectime"
	"danielromeroauk/golang-demo/pkg/number"
	"danielromeroauk/golang-demo/pkg/sum"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sumWithGoroutines(numbers []int, goroutines int) int {
	sliceSize := len(numbers) / goroutines

	// tool for avoid race conditions
	mu := &sync.RWMutex{}

	// tool for wait for all goroutines
	wg := &sync.WaitGroup{}

	total_sum := 0
	for i := 0; i < goroutines; i++ {
		var start, end int
		if i < goroutines-1 {
			start = i * sliceSize
			end = (i + 1) * sliceSize
		} else {
			start = i * sliceSize
			end = len(numbers)
		}

		wg.Add(1)
		go func(elems []int) {
			defer wg.Done()
			sum := sum.Integers(elems)

			// avoid race conditions
			mu.Lock()
			total_sum += sum
			mu.Unlock()
		}(numbers[start:end])
	}

	// wait for all goroutines
	wg.Wait()

	return total_sum
}

func main() {
	runtime.GOMAXPROCS(4)
	defer exectime.Log(time.Now(), "sum")

	numbers := number.GetSliceInt(number.Billion)

	// a.k.a. workers
	const goroutines int = 8

	total_sum := sumWithGoroutines(numbers, goroutines)
	fmt.Println(total_sum)
}

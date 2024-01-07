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

func sumSubslice(subslice []int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- sum.Integers(subslice)
}

func sumWithGoroutines(numbers []int, goroutines int) int {
	if goroutines <= 0 {
		goroutines = 1
	}

	total := 0
	ch := make(chan int)
	var wg sync.WaitGroup

	sliceSize := len(numbers) / goroutines

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
		go sumSubslice(numbers[start:end], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for partialSum := range ch {
		total += partialSum
	}

	return total
}

func main() {
	runtime.GOMAXPROCS(4)
	defer exectime.Log(time.Now(), "sum")

	numbers := number.GetSliceInt(number.Billion)
	goroutines := 8

	total_sum := sumWithGoroutines(numbers, goroutines)
	fmt.Println(total_sum)
}

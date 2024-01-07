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

func main() {
	runtime.GOMAXPROCS(2)
	defer exectime.Log(time.Now(), "sum")

	numbers := number.GetSliceInt(number.Billion)

	// tool for avoid race conditions
	mu := &sync.RWMutex{}

	// tool for wait for all goroutines
	wg := &sync.WaitGroup{}

	total_sum := 0

	group1 := numbers[:len(numbers)/2]
	group2 := numbers[len(numbers)/2:]

	wg.Add(2)
	go GoSum(group1, wg, mu, &total_sum)
	go GoSum(group2, wg, mu, &total_sum)

	// wait for all goroutines
	wg.Wait()

	fmt.Println(total_sum)
}

func GoSum(numbers []int, wg *sync.WaitGroup, mu *sync.RWMutex, total_sum *int) {
	defer wg.Done()
	sum := sum.Integers(numbers)

	// avoid race conditions
	mu.Lock()
	*total_sum += sum
	mu.Unlock()
}

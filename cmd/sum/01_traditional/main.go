package main

import (
	"danielromeroauk/golang-demo/pkg/exectime"
	"danielromeroauk/golang-demo/pkg/number"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	defer exectime.Log(time.Now(), "sum")

	numbers := number.GetSliceInt(number.Billion)
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}

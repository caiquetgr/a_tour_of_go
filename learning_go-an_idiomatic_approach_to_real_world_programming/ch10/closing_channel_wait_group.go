package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	concurrency := 10
	in := make(chan int, concurrency)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			in <- rand.Intn(10)
		}()
	}

	go func() {
		wg.Wait()
		close(in)
	}()

	results := processAndGather(
		in,
		doubleSize,
		concurrency,
	)

	fmt.Println(results)
}

func doubleSize(i int) int {
	return i * 2
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)

	var wg sync.WaitGroup
	wg.Add(num)

	for v := range in {
		go func(val int) {
			defer wg.Done()
			out <- processor(val)
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int

	for v := range out {
		result = append(result, v)
	}

	return result
}

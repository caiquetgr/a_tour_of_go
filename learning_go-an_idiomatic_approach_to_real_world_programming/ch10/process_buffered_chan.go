package main

import "fmt"

func main() {
	ch := make(chan int, 20)

	for i := 0; i < 20; i++ {
		ch <- i
	}

  out := processChannel(ch)

  fmt.Println(out)
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)

	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}

	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	return v * 2
}

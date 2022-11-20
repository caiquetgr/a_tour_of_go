package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))

	for _, v := range a {
		go func() {
			// bug - using for variable
			ch <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Correcting the bug shadowing the value:")

	// shadowing the value
	for _, v := range a {
		v := v
		go func() {
			ch <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}

  fmt.Println("Correcting the bug passing variable to couroutine:")
	// variable to coroutine
	for _, v := range a {
		go func(val int) {
			ch <- val * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

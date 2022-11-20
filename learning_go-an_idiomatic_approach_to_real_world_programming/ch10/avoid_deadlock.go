package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		fmt.Println("wrote v on ch1:", v)
		v2 := <-ch2
		fmt.Println("read v2 on ch2:", v2)
		fmt.Println(v, v2)
	}()

	v := 2
	var v2 int

	// select avoids deadlock
	select {
	case ch2 <- v:
		fmt.Println("wrote v on ch2:", v)
	case v2 = <-ch1:
		fmt.Println("read v2 from ch1:", v2)
	// default:
		// fmt.Println("no value written to ch")
	}

	fmt.Println(v, v2)

}

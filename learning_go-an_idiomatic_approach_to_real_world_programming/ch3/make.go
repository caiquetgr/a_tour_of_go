package main

import "fmt"

func main() {

	x := make([]int, 0, 10) // length = 5, capacity = 10
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x, len(x), cap(x))
}

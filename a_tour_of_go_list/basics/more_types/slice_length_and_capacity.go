package main

import "fmt"

func main() {
	s := []int{2, 35, 7, 11, 13}
	printSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// extend it's length
	s = s[:4]
	printSlice(s)

	// drop it's first two values
	s = s[:2]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

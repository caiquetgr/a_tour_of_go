package main

import "fmt"

func main() {
	/*
		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		num = copy(y, x)

		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		copy(y, x[2:])

		x := []int{1, 2, 3, 4}
		num = copy(x[:3], x[1:])
		fmt.Println(x, num)
	*/

	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}

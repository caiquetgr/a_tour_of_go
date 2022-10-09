package main

import "fmt"

func main() {
	//var x = []int{10, 20, 30}
	//var x = []int{1, 5:4, 6}
	//var x [][]int
	var x []int
	fmt.Println(x == nil) //true

	x = append(x, 10)
	x = append(x, 5, 6, 7)

	y := []int{20, 30, 40}
	x = append(x, y...)

	a := []int{1, 2, 3, 4}
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	e := a[:]
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

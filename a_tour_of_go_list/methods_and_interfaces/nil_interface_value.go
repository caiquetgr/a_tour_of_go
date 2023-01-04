package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // runtime error (doesnt know which concrete M to call)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

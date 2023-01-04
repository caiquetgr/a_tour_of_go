
package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

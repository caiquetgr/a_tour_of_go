package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // if i is not string, it panics
	fmt.Println(s)

	s, ok := i.(string) // avoiding panic
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

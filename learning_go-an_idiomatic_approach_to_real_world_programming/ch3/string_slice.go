package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(s, b)

	var s2 = s[4:7]
	fmt.Println(s2)

	s = "Hello 🌞"
	fmt.Println(s[6:])
	fmt.Println(len(s)) // len is 10, because of sun emoji is utf8 four bytes. string are counted as one char = one byte utf8
}

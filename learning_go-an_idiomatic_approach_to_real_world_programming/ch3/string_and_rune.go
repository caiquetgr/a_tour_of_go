package main

import "fmt"

func main() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(a, s, b, s2)

	var x int = 65
	var y = string(x)
	fmt.Println(y) // prints A

	var s3 string = "Hello, 🌞"
	var bs []byte = []byte(s3)
	var rs []rune = []rune(s3)
	fmt.Println(bs)
	fmt.Println(rs)
}

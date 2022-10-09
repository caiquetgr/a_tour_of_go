package main

import "fmt"

func main() {
	//var x [3]int
	//var x = [3]int{10, 20, 30}
	//var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	//var x = [...]int{10, 20, 30}
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
	fmt.Println(x[1])
	fmt.Println(len(x))

	//var x [2][3]int
}

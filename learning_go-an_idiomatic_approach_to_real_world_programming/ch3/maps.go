package main

import "fmt"

func main() {
	//var nilMap map[string]int // nilMap, read from int returns zero, write in it causes panic
	//totalWins := map[string]int{} // not nil map, can read and write form it

	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Bilie"},
	}
	fmt.Println(teams)

	ages := make(map[int][]string, 10) // length 0, and grows past the initially specified size
	fmt.Println(ages)

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")
}

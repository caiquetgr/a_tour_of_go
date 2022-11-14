package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once

	for i := 0; i < 3; i++ {
		printName("Teste 1")
		once.Do(func() {
			printName("Teste 2")
		})
	}

}

func printName(name string) {
	fmt.Println(name)
}

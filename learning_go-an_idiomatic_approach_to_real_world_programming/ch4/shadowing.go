package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// there is a tool to detect shadowing: go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
}

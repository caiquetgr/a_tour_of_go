package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)

	m := make(map[string]int)

	for _, v := range tokens {
		_, isPresent := m[v]
		if isPresent {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}

package main

import (
	"fmt"
	"strings"
)

type TextHandler interface {
	process(text string)
}

type TextHandlerFunc func(string)

func (a TextHandlerFunc) process(text string) {
	a(text)
}

func main() {
	processText("oi", TextHandlerFunc(func(s string) { fmt.Println(strings.ToUpper(s)) }))
	processText("hello ", TextHandlerFunc(printTextFiveTimes))
	processText2("oi", func(s string) { fmt.Println(strings.ToUpper(s)) })
	processText2("hello ", printTextFiveTimes)
}

func printTextFiveTimes(text string) {
	for i := 0; i < 5; i++ {
		fmt.Print(text)
	}
}

func processText(text string, handler TextHandler) {
	handler.process(text)
}

func processText2(text string, f func(string)) {
	f(text)
}

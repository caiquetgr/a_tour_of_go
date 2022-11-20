package main

import (
	"fmt"
	"net/http"
)

func main() {

	ch := make(chan string)

	links := []string{
		"http://abc.com",
		"http://pqr.com",
		"http://zyx.com",
	}

	for _, link := range links {
		go makeHttpCall(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Printf("Link %v is up\n", <-ch)
	}
}

func makeHttpCall(link string, ch chan string) {
  _, err := http.Get(link)
  if err == nil {
    ch <- link
  }
}

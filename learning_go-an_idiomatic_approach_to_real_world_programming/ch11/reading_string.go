package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "Testing a word counter! It looks like with works!!"
	sr := strings.NewReader(s)

	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println("oh noes")
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buff := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buff)

		for _, b := range buff[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}

	}
}

package main

import (
	"compress/gzip"
	"io"
	"os"
)

func process(r io.Reader) error {
}

func main() {
	r, err := os.Open("testfile")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	gz, err := gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	defer gz.Close()
	process(gz)
}

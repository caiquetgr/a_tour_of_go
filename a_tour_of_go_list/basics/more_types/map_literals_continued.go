package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {20.123, 29.123},
	"Google":    {20.111, 20.222},
}

func main() {
	fmt.Println(m)
}

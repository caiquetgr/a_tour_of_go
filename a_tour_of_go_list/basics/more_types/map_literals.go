package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.4897, 98.8777,
	},
	"Google": Vertex{
		12.4242, 90.998,
	},
}

func main() {
	fmt.Println(m)
}

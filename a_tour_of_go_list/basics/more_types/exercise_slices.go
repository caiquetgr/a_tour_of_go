package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, 0, dy)
	for i := 0; i < dx; i++ {
		x = append(x, make([]uint8, dx))
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			x[i][j] = uint8(i ^ j)
		}
	}

	return x
}

func main() {
	pic.Show(Pic)
}

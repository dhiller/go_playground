package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := range y {
		y[i] = make([]uint8, dx)
		for k := range y[i] {
			y[i][k] = uint8(i * k)
		}
	}
	return y
}

func main() {
	pic.Show(Pic)
}


package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for math.Abs(z * z - x) > 0.00000000001 {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrt(float64(i)))
	}
}

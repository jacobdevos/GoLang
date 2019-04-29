package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Pow(z, 2) != x {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	num := 49.0
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}

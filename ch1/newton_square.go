package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var newZ float64
	for x-z >= 0.1 {

		newZ = z - ((math.Pow(z, 2) - x) / (2 * z))
		if (newZ - z) <= 0.1 {
			return newZ
		} else {
			z = newZ
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(25))
}

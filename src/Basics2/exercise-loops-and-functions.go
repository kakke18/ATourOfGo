package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for v := 0; v < 10; v++ {
		z -= (z * z - x) / (2.0 * z)
		fmt.Printf("z:%g\n", z)
	}
	return z
}

func main() {
	fmt.Println("\nresult:", Sqrt(2.0))
	fmt.Println("answer:", math.Sqrt(2.0))
}

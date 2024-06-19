package main

import (
	"fmt";
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	eps := float64(1e-5)
	change := float64(1)
	num_iter := 1

	for change > eps || -change > eps {
		change = (z*z - x) / (2*z)
		z -= change
		num_iter += 1
	}
	
	fmt.Printf("The approximated square root value is %.5f\n", z)
	fmt.Printf("The number of iterations is %d\n", num_iter)
	fmt.Printf("math.Sqrt result: %.5f\n", math.Sqrt(x))
	return z
}

func main() {
	fmt.Println(Sqrt(288))
}

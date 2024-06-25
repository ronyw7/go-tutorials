package main

import (
	"fmt";
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	eps := float64(1e-5)
	change := float64(1)
	num_iter := 1

	for change > eps || -change > eps {
		change = (z*z - x) / (2*z)
		z -= change
		num_iter += 1
	}
	return z, nil
}

func main() {
	values := []float64{237, -5, 2, -16}
	for _, value := range values {
		result, err := Sqrt(value)
		if err != nil {
			fmt.Printf("\n")
			fmt.Println(err)
		} else {
			fmt.Printf("The square root of %v is %.5f\n", value, result)
			fmt.Printf("math.Sqrt result: %.5f\n", math.Sqrt(value))
		}
	}
}

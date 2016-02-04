package main

import (
	"math"
	"fmt"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

func sum(i, j int) (sum int) {
	sum = i + j
	return sum
}

func main() {
	fmt.Printf("%f\n", hypot(2.0, 2.0))

	theSum := sum(1,2)
	fmt.Printf("theSum: %d\n", theSum)
}


package main

import (
	"math"
	"fmt"
	"strings"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

func sum(i, j int) (sum int) {
	sum = i + j
	return sum
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func add1(r rune) rune {
	return r + 1
}

func main() {
	fmt.Printf("%f\n", hypot(2.0, 2.0))

	theSum := sum(1, 2)
	fmt.Printf("theSum: %d\n", theSum)

	f := square
	fmt.Printf("square(%d) is %d\n", 3, f(3))

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}


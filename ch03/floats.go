package main

import (
	"fmt"
	"math"
)

// 3.2 Floats
func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	const e = 2.71828 // (approximately)
	const Avogadro = 6.02214129e23
	const Planck   = 6.62606957e-34

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"
}

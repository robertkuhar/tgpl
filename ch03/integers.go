package main

import "fmt"

// 3.1 Integers
func main() {
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges // compile error
	var compote int = int(apples) + int(oranges) // compile error
	fmt.Printf("%d\n", compote)

	var ima32 int32 = 32
	var ima16 int16 = 32
	if int(ima32) == int(ima16) {
		fmt.Printf("%d is equal to %d\n", ima32, ima16)
	} else {
		fmt.Printf("%d is NOT equal to %d\n", ima32, ima16)
	}

	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i)   // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	f1 := 1e100  // a float64
	i1 := int(f1) // result is implementation-dependent
	fmt.Printf("f1 is %e, i1 is %d\n", f1, i1)

	f2 := 1e3
	i2 := int(f2)
	fmt.Printf("f2 is %e, i2 is %d\n", f2, i2)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}

package main
import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...) // append the slice z
	fmt.Printf("z: %v\n", z)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"

	var z1 []int
	z1 = appendInt2(z1, 1)
	z1 = appendInt2(z1, 2, 3)
	z1 = appendInt2(z1, 4, 5, 6)
	z1 = appendInt2(z1, z1...)
	fmt.Printf("z1: %v\n", z1)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// ...expand z to at least zlen...
	copy(z[len(x):], y)
	return z
}


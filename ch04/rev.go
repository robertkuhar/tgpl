package main
import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Printf("a: %v\n", a) // "[5 4 3 2 1 0]"

	// Rotate s left by two positions.
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Printf("a: %v\n", a) // "[3 2 1 0 5 4]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

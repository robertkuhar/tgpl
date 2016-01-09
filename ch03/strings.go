package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')

	// 104?  WTF?  Same with the rune.  What?
	c := 'h'
	fmt.Println(c)

	fmt.Println(s[0:5])

	// This panics us with index out of range
	//	c1 := s[len(s)]
	//	fmt.Println(c1)

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	s1 := "left foot"
	t := s1
	s1 += ", right foot"
	fmt.Println(s1) // "left foot, right foot"
	fmt.Println(t) // "left foot"
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func lenOfString(label string, s string) {
	fmt.Printf("%s's value is \"%s\", len is %d, utf8.RuneCountInString is %d\n",
		label,
		s,
		len(s),
		utf8.RuneCountInString(s))
}

func printRunesInString(label string, s string) {
	fmt.Printf("%s printRunesInString(\"%s\")\n", label, s)
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

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

	helloNonAscii := "Hello, 世界"
	lenOfString("helloNonAscii", helloNonAscii)
	printRunesInString("helloNonAscii",helloNonAscii)
	helloAscii := "Hello, Robert"
	lenOfString("helloAscii", helloAscii)
	printRunesInString("helloAscii",helloAscii)

}

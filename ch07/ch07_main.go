package main

import (
	"fmt"
	"bytes"
	"github.com/robertkuhar/tgpl/ch07/bytecounter"
)

func main() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "firstName: %s, lastName: %s", "robert", "kuhar")
	s := string(buf.Bytes())
	fmt.Printf("string from buf.Bytes(): %s\n", s)
	fmt.Printf("string from buf.String(): %s\n", buf.String())

	// 7.1 gopl.io/ch7/bytecounter
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")


}

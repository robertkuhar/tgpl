// calls popcount
package main

import (
	"github.com/robertkuhar/tgpl/ch02/popcount" 
	"flag" 
	"fmt"
)

// TODO: This is just unnatural to me, why does this return a pointer?
var x = flag.Uint64("x", 0, "x is an unsigned integer")

func main() {
	flag.Parse()
        bits := popcount.PopCount(*x)
	fmt.Printf("%d has %d bits set\n",*x,bits)
}


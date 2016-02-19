package main

import (
	"fmt"
	"github.com/robertkuhar/tgpl/mypackage"
)

func main() {
	x := 2
	y := 4
	fmt.Printf("Add(x:%d, y:%d) %d\n", x, y, mypackage.Add(x, y))
	fmt.Printf("Sub(x:%d, y:%d) %d\n", x, y, mypackage.Sub(x, y))
}

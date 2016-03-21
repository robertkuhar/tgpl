package main

import (
	"github.com/robertkuhar/tgpl/mypackage"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", mypackage.Add(1,1));
	fmt.Printf("%d\n", mypackage.Sub(1,1));

	var imma_int int = 1
	if true {
		imma_int := "Jello"
		fmt.Printf("in if, imma_int is %s\n", imma_int)
	}
	fmt.Printf("outside of if, imma_int is %d\n", imma_int)

}

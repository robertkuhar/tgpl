package main

import (
	"flag" 
	"fmt"
)

var intFromFlag = flag.Int("int", 0, "int here")
var int64FromFlag = flag.Int64("int64", 0, "int64 here")

func main() {
	flag.Parse()
	fmt.Printf("intFromFlag\t%d\n",*intFromFlag)
	fmt.Printf("int64FromFlag\t%d\n",*int64FromFlag)
}

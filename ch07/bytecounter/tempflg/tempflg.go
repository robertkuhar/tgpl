package main

import (
	"flag"
	"fmt"
	"github.com/robertkuhar/tgpl/ch07/tempconv"
)

// It is interesting that the Flag doesn't have the usage string, the client does.
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

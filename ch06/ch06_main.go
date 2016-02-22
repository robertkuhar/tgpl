package main

import (
	"github.com/robertkuhar/tgpl/ch06/geometry"
	"fmt"
)

func main() {
	perim := geometry.Path {
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Printf("distance: %f, cost: %f\n", perim.Distance(), perim.Cost())


	r := &geometry.Point{1,2}
	r.ScaleBy(2)
	fmt.Println(*r)
}

package coloredpoint

import "image/color"

// Lookie here, the oldest form of inheritance.  Cut-and-Past from geometry package
type Point struct { X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}
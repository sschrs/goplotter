package goplotter

import "image/color"

type Point struct {
	X, Y       int
	pointShape Shape
}

func NewPoint(x, y int, shape Shape) Point {
	return Point{
		X:          x,
		Y:          y,
		pointShape: shape,
	}
}

func SquarePoint(x, y, size int, clr color.Color) Point {
	return Point{
		X:          x,
		Y:          y,
		pointShape: NewRectangle(size, size, clr),
	}
}

func CirclePoint(x, y, r int, clr color.Color) Point {
	return Point{
		X:          x,
		Y:          y,
		pointShape: NewCircle(r, clr),
	}
}

func TrianglePoint(x, y, size int, clr color.Color) Point {
	return Point{
		X:          x,
		Y:          y,
		pointShape: NewTriangle(size, size, clr),
	}
}

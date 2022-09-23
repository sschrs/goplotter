package goplotter

type Point struct {
	X, Y       int
	pointShape Shape
}

func NewPoint(x, y, size int, shape Shape) *Point {
	return &Point{
		X:          x,
		Y:          y,
		pointShape: shape,
	}
}

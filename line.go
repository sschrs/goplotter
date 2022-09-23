package goplotter

import "image/color"

type Line struct {
	StartPoint, EndPoint, Width int
	Color                       color.Color
}

func NewLine(from, to, width int, clr color.Color) Line {
	return Line{
		StartPoint: from,
		EndPoint:   to,
		Width:      width,
		Color:      clr,
	}
}

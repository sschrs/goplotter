package goplotter

import "image/color"

type Line struct {
	StartX, StartY, EndX, EndY, Width int
	Color                             color.Color
}

func NewLine(fromX, fromY, toX, toY, width int, clr color.Color) Line {
	return Line{
		StartX: fromX,
		StartY: fromY,
		EndX:   toX,
		EndY:   toY,
		Width:  width,
		Color:  clr,
	}
}

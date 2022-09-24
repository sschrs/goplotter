package goplotter

import "image/color"

type Axis struct {
	Labels []string
	Width  int
	Color  color.Color
}

func NewAxis(labels []string, width int, clr color.Color) Axis {
	return Axis{
		Labels: labels,
		Width:  width,
		Color:  clr,
	}
}

package goplotter

import "image/color"

type Axis struct {
	Labels []*Label
	Width  int
	Color  color.Color
}

func NewAxis(labels []*Label, width int, clr color.Color) *Axis {
	return &Axis{
		Labels: labels,
		Width:  width,
		Color:  clr,
	}
}

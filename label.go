package goplotter

import "image/color"

type Label struct {
	Name     string
	Color    color.Color
	Position int
}

func NewLabel(name string, clr color.Color) *Label {
	return &Label{Name: name, Color: clr}
}

func (label *Label) SetPosition(position int) {
	label.Position = position
}

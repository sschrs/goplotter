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

func NewLabels(names []string, clr color.Color) []*Label {
	var labels []*Label
	for _, name := range names {
		labels = append(labels, &Label{Name: name, Color: clr})
	}
	return labels
}

func (label *Label) SetPosition(position int) {
	label.Position = position
}

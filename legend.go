package goplotter

import "image/color"

var (
	defaultBackgroundColor = color.White
	defaultTitleColor      = color.Black
	defaultTextColor       = color.Black
)

type LegendItem struct {
	Text      string
	TextColor color.Color
	Symbol    Shape
}

type Legend struct {
	Title                       string
	TitleColor, BackgroundColor color.Color
	Items                       []LegendItem
	Width, Height, X, Y         int
}

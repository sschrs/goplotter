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
	Items                       []*LegendItem
	Width, Height, X, Y         int
}

func NewLegendItem(text string, symbol Shape) *LegendItem {
	return &LegendItem{
		Text:      text,
		TextColor: defaultTextColor,
		Symbol:    symbol,
	}
}

func NewLegend(title string, items []*LegendItem, width, height, x, y int) *Legend {
	return &Legend{
		Title:           title,
		TitleColor:      defaultTitleColor,
		BackgroundColor: defaultBackgroundColor,
		Items:           items,
		Width:           width,
		Height:          height,
		X:               x,
		Y:               y,
	}
}

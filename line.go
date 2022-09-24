package goplotter

import (
	"image"
	"image/color"
	"math"
)

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

func bresenhamLine(x0, x1, y0, y1 int, img *image.RGBA, clr color.Color) *image.RGBA {
	x := x0
	y := y0
	dx := int(math.Abs(float64(x1-x0)) + 0.5)
	dy := int(math.Abs(float64(y1-y0)) + 0.5)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}

	interchane := 0

	if dy > dx {
		temp := dy
		dy = dx
		dx = temp
		interchane = 1
	}
	err := 2*dy - dx

	for i := 0; i < dx-1; i++ {
		if err > 0 {
			if interchane == 1 {
				x += sx
			} else {
				y += sy
			}
			err -= 2 * dx
		}
		if interchane == 1 {
			y += sy
		} else {
			x += sx
		}
		err += 2 * dy
		img.Set(x, y, clr)
	}
	return img
}

package goplotter

import (
	"image"
	"image/color"
)

type Shape image.Image

func NewCircle(r int, shapeColor color.Color) Shape {
	img := image.NewRGBA(image.Rect(0, 0, r*2, r*2))
	x0 := r
	y0 := r
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x >= y {
		img.Set(x0+x, y0+y, shapeColor)
		img.Set(x0+y, y0+x, shapeColor)
		img.Set(x0-y, y0+x, shapeColor)
		img.Set(x0-x, y0+y, shapeColor)
		img.Set(x0-x, y0-y, shapeColor)
		img.Set(x0-y, y0-x, shapeColor)
		img.Set(x0+y, y0-x, shapeColor)
		img.Set(x0+x, y0-y, shapeColor)

		for i := x0 - x; i <= x0+x; i++ {
			img.Set(i, y0+y, shapeColor)
		}

		for i := x0 - y; i <= x0+y; i++ {
			img.Set(i, y0+x, shapeColor)
		}

		for i := x0 - x; i < x0+x; i++ {
			img.Set(i, y0-y, shapeColor)
		}

		for i := x0 - y; i < x0+y; i++ {
			img.Set(i, y0-x, shapeColor)
		}

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
	return img
}

func NewRectangle(width, height int, shapeColor color.Color) Shape {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			img.Set(i, j, shapeColor)
		}
	}

	return img
}

func NewTriangle(baseWidth, height int, shapeColor color.Color) Shape {
	img := image.NewRGBA(image.Rect(0, 0, baseWidth, height))

	positionX := int(baseWidth / 2)
	positionY := 0

	for i := 1; i <= height/2; i++ {
		positionX -= 1
		positionY += 1
		img.Set(positionX, positionY, shapeColor)
		for j := 0; j < i*2; j++ {
			img.Set(positionX+j, positionY, shapeColor)
		}

	}
	return img
}

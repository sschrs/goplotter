package goplotter

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Plot struct {
	Width, Height         int
	Title, XLabel, YLabel string
	BackroundColor        color.Color
}

func NewPlot(width, height int) *Plot {
	return &Plot{
		Width:          width,
		Height:         height,
		BackroundColor: color.White,
	}
}

func (plot Plot) Draw(path string) error {
	upLeft := image.Point{X: 0, Y: 0}
	lowReight := image.Point{X: plot.Width, Y: plot.Height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowReight})

	for x := 0; x < plot.Width; x++ {
		for y := 0; y < plot.Height; y++ {
			img.Set(x, y, plot.BackroundColor)
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	png.Encode(f, img)

	return nil
}

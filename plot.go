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
	Points                []Point
	Lines                 []Line
	AxisX, AxisY          Axis
}

func NewPlot(width, height int, axisX, axisY Axis) *Plot {
	return &Plot{
		Width:          width,
		Height:         height,
		BackroundColor: color.White,
		AxisX:          axisX,
		AxisY:          axisY,
	}
}

func (plot *Plot) Draw(path string) error {
	upLeft := image.Point{X: 0, Y: 0}
	lowReight := image.Point{X: plot.Width, Y: plot.Height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowReight})

	for x := 0; x < plot.Width; x++ {
		for y := 0; y < plot.Height; y++ {
			img.Set(x, y, plot.BackroundColor)
		}
	}
	plot.plotAxes()
	img = plot.plotLines(img)
	img = plot.plotPoints(img)

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	png.Encode(f, img)

	return nil
}

func (plot *Plot) AddPoint(point Point) {
	plot.Points = append(plot.Points, point)
}

func (plot *Plot) AddPoints(points []Point) {
	for _, point := range points {
		plot.Points = append(plot.Points, point)
	}
}

func (plot *Plot) AddLine(line Line) {
	plot.Lines = append(plot.Lines, line)
}

func (plot *Plot) AddLines(lines []Line) {
	for _, line := range lines {
		plot.Lines = append(plot.Lines, line)
	}
}

func (plot *Plot) plotAxes() {
	padding := 50
	xLine := NewLine(padding, plot.Height-padding, plot.Width-padding, plot.Height-padding, plot.AxisY.Width, plot.AxisY.Color)
	yLine := NewLine(padding, padding, padding, plot.Height-padding, plot.AxisX.Width, plot.AxisX.Color)
	plot.AddLines([]Line{xLine, yLine})

	if len(plot.AxisX.Labels) > 0 {
		xInterval := (xLine.EndX - xLine.StartX) / len(plot.AxisX.Labels)
		y := plot.Height - padding
		for i := xLine.StartX; i <= xLine.EndX; i += xInterval {
			plot.AddLine(NewLine(i, y-5, i, y+5, plot.AxisX.Width, plot.AxisX.Color))
		}

	}

	if len(plot.AxisY.Labels) > 0 {
		yInterval := (yLine.EndY - yLine.StartY) / len(plot.AxisY.Labels)
		x := padding
		for i := yLine.StartY; i <= yLine.EndY; i += yInterval {
			plot.AddLine(NewLine(x-5, i, x+5, i, plot.AxisY.Width, plot.AxisY.Color))
		}

	}

}

func (plot *Plot) plotPoints(img *image.RGBA) *image.RGBA {
	for _, point := range plot.Points {
		point.X -= point.pointShape.Bounds().Size().X / 2
		point.Y -= point.pointShape.Bounds().Size().Y / 2
		for x := point.X; x <= point.X+point.pointShape.Bounds().Size().X; x++ {
			for y := point.Y; y <= point.Y+point.pointShape.Bounds().Size().Y; y++ {
				r, g, b, a := point.pointShape.At(x-point.X, y-point.Y).RGBA()
				if r != 0 || g != 0 || b != 0 || a != 0 {
					img.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
				}
			}
		}
	}
	return img
}

func (plot *Plot) plotLines(img *image.RGBA) *image.RGBA {
	for _, line := range plot.Lines {
		slope := float64(line.EndY-line.StartY) / float64(line.EndX-line.StartX)
		if slope > 1 {
			x0, x1, y0, y1 := line.StartX-line.Width/2, line.EndX-line.Width/2, line.StartY, line.EndY
			for i := 0; i < line.Width; i++ {
				img = bresenhamLine(x0, x1, y0, y1, img, line.Color)
				x0++
				x1++
			}
		} else {
			x0, x1, y0, y1 := line.StartX, line.EndX, line.StartY-line.Width/2, line.EndY-line.Width/2
			for i := 0; i < line.Width; i++ {
				img = bresenhamLine(x0, x1, y0, y1, img, line.Color)
				y0++
				y1++
			}
		}

	}
	return img
}

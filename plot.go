package goplotter

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

const (
	fontWidth  = 8
	fontHeight = 16
)

type Plot struct {
	Width, Height                                         int
	Title, XLabel, YLabel                                 string
	BackgroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	Points                                                []Point
	Lines                                                 []Line
	AxisX, AxisY                                          *Axis
	Legend                                                Legend
}

func NewPlot(width, height int, axisX, axisY *Axis) *Plot {
	return &Plot{
		Width:           width,
		Height:          height,
		BackgroundColor: color.White,
		AxisX:           axisX,
		AxisY:           axisY,
		TitleColor:      color.Black,
		XLabelColor:     color.Black,
		YLabelColor:     color.Black,
	}
}

func (plot *Plot) Draw() *image.RGBA {
	// Create an empty image
	upLeft := image.Point{X: 0, Y: 0}
	lowReight := image.Point{X: plot.Width, Y: plot.Height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowReight})

	// Fill the image with background color
	for x := 0; x < plot.Width; x++ {
		for y := 0; y < plot.Height; y++ {
			img.Set(x, y, plot.BackgroundColor)
		}
	}

	if plot.Title != "" {
		titleWidth := len(plot.Title) * fontWidth
		firstPoint := (plot.Width / 2) - (titleWidth / 2)
		plot.plotText(img, firstPoint, 40, plot.Title, plot.TitleColor)
	}

	plot.plotAxes(img)
	img = plot.plotLines(img)
	img = plot.plotPoints(img)
	plot.plotLegend(img)

	return img
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

func (plot *Plot) plotAxes(img *image.RGBA) {
	padding := 70
	xLine := NewLine(padding, plot.Height-padding, plot.Width-padding, plot.Height-padding, plot.AxisY.Width, plot.AxisY.Color)
	yLine := NewLine(padding, padding, padding, plot.Height-padding, plot.AxisX.Width, plot.AxisX.Color)
	plot.AddLines([]Line{xLine, yLine})

	if len(plot.AxisX.Labels) > 0 {
		xInterval := (xLine.EndX - xLine.StartX) / (len(plot.AxisX.Labels) - 1)
		y := plot.Height - padding
		labelx := 0
		for i := xLine.StartX; i <= xLine.EndX; i += xInterval {
			plot.AddLine(NewLine(i, y-5, i, y+5, plot.AxisX.Width, plot.AxisX.Color))

			text := plot.AxisX.Labels[labelx].Name
			textLength := len(text)
			X := i - (fontWidth / 2 * textLength)
			plot.plotText(img, X, y+20, plot.AxisX.Labels[labelx].Name, plot.AxisX.Labels[labelx].Color)
			plot.AxisX.Labels[labelx].SetPosition(i)

			labelx++
		}

	}

	labely := len(plot.AxisY.Labels) - 1
	if len(plot.AxisY.Labels) > 0 {
		yInterval := (yLine.EndY - yLine.StartY) / (len(plot.AxisY.Labels) - 1)
		x := padding
		for i := yLine.StartY; i <= yLine.EndY; i += yInterval {
			plot.AddLine(NewLine(x-5, i, x+5, i, plot.AxisY.Width, plot.AxisY.Color))
			if labely != -1 {
				text := plot.AxisY.Labels[labely].Name
				textWidth := len(text) * fontWidth
				plot.plotText(img, x-textWidth-5, i, text, plot.AxisY.Labels[labely].Color)
				plot.AxisY.Labels[labely].Position = i
			}
			labely--
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

func (plot *Plot) plotText(img *image.RGBA, x, y int, text string, clr color.Color) {
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(clr),
		Face: inconsolata.Bold8x16,
		Dot:  point,
	}

	d.DrawString(text)
}

func (plot *Plot) plotLegend(img *image.RGBA) {
	if len(plot.Legend.Items) > 0 {
		// Create legend rectangle with background color
		for x := plot.Legend.X; x <= plot.Legend.Width+plot.Legend.X; x++ {
			for y := plot.Legend.Y; y <= plot.Legend.Height+plot.Legend.Y; y++ {
				img.Set(x, y, plot.Legend.BackgroundColor)
			}
		}

		// Add legend title
		x := (plot.Legend.X + plot.Legend.Width/2) - (len(plot.Legend.Title)*fontWidth)/2
		y := plot.Legend.Y + 30
		plot.plotText(img, x, y, plot.Legend.Title, plot.Legend.TitleColor)

		// Add Legend Items
		x = plot.Legend.X + 20
		y = plot.Legend.Y + 60
		for _, item := range plot.Legend.Items {
			point := NewPoint(x, y, item.Symbol)
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
			plot.plotText(img, x+point.pointShape.Bounds().Size().X+5, y+point.pointShape.Bounds().Size().Y/2-4, item.Text, item.TextColor)
			y += 30
		}

	}
}

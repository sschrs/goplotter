package plotter

import (
	"fmt"
	"github.com/sschrs/goplotter"
	"image"
	"image/color"
	"strconv"
)

type ScatterPlotter struct {
	XValues, YValues                                      []float64
	Title, XLabel, YLabel                                 string
	BackgroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	Legend                                                bool
	Width, Height                                         int
	PointShape                                            goplotter.Shape
	XRange, YRange                                        Range
}

func NewScatterPlot(x, y []float64) *ScatterPlotter {
	return &ScatterPlotter{
		XValues:         x,
		YValues:         y,
		XLabelColor:     color.Black,
		YLabelColor:     color.Black,
		BackgroundColor: color.White,
		TitleColor:      color.Black,
		Legend:          false,
		Width:           1280,
		Height:          720,
		PointShape:      goplotter.NewCircle(5, color.RGBA{R: 255, A: 255}),
	}
}

func (sp *ScatterPlotter) Plot() *image.RGBA {
	// Set XAxis Range
	_, maxX := minMax(sp.XValues)
	if sp.XRange.Step == 0 {
		step := (maxX) / 10
		sp.XRange.From = 0
		sp.XRange.To = maxX
		sp.XRange.Step = step
	}

	// Set YAxis Range
	_, maxY := minMax(sp.YValues)
	if sp.YRange.Step == 0 {
		step := (maxY) / 10
		sp.YRange.From = 0
		sp.YRange.To = maxY
		sp.YRange.Step = step
	}

	// Set Axis Labels
	var xLabelValues, yLabelValues []string

	// X Axis Labels
	for i := sp.XRange.From; i <= sp.XRange.To; i += sp.XRange.Step {
		i, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", i), 64)
		xLabelValues = append(xLabelValues, fmt.Sprintf("%.2f", i))
	}
	// Y Axis Labels
	for i := sp.YRange.From; i <= sp.YRange.To; i += sp.YRange.Step {
		i, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", i), 64)
		yLabelValues = append(yLabelValues, fmt.Sprintf("%.2f", i))
	}

	xLabels := goplotter.NewLabels(xLabelValues, sp.XLabelColor)
	yLabels := goplotter.NewLabels(yLabelValues, sp.YLabelColor)

	axisX := goplotter.NewAxis(xLabels, 3, sp.XLabelColor)
	axisY := goplotter.NewAxis(yLabels, 3, sp.YLabelColor)

	// Create Plot
	plot := goplotter.NewPlot(sp.Width, sp.Height, axisX, axisY)
	plot.BackgroundColor = sp.BackgroundColor
	plot.Title = sp.Title
	plot.TitleColor = sp.TitleColor

	plot.Draw()

	// Add Points
	for i := 0; i < len(sp.XValues); i++ {
		x := sp.XValues[i]
		y := sp.YValues[i]

		endPointX := sp.Width - 70
		startPointX := 70
		lengthX := endPointX - startPointX
		xPosition := (float64(lengthX)*x)/float64(maxX) + 70

		startPointY := sp.Height - 70
		endPointY := 70
		lengthY := startPointY - endPointY
		yPosition := float64(sp.Height) - float64(lengthY)*y/float64(maxY) - 70

		plot.AddPoint(goplotter.NewPoint(int(xPosition+0.5), int(yPosition+0.5), sp.PointShape))
	}

	return plot.Draw()
}

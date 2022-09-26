package plotter

import (
	"fmt"
	"github.com/sschrs/goplotter"
	"image"
	"image/color"
	"sort"
)

type ScatterPlotter struct {
	XValues, YValues                                      []string
	Title, XLabel, YLabel                                 string
	BackgroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	AddLegend                                             bool
	Width, Height                                         int
	PointShape                                            goplotter.Shape
}

func NewScatterPlot(x []any, y []float64) *ScatterPlotter {
	var xLabels []string
	var yLabels []string

	for _, xValue := range x {
		xLabels = append(xLabels, fmt.Sprintf("%v", xValue))
	}

	for _, yValue := range y {
		yLabels = append(yLabels, fmt.Sprintf("%.2f", yValue))
	}

	return &ScatterPlotter{
		XValues:         xLabels,
		YValues:         yLabels,
		XLabelColor:     color.Black,
		YLabelColor:     color.Black,
		BackgroundColor: color.White,
		TitleColor:      color.Black,
		AddLegend:       false,
		Width:           1280,
		Height:          720,
		PointShape:      goplotter.NewCircle(5, color.RGBA{R: 255, A: 255}),
	}
}

func (sp *ScatterPlotter) Plot() *image.RGBA {
	xLabels := goplotter.NewLabels(sp.XValues, sp.XLabelColor)
	var yAxisValues []string

	for _, v := range sp.YValues {
		yAxisValues = append(yAxisValues, v)
	}

	sort.Slice(yAxisValues, func(i, j int) bool {
		return yAxisValues[i] < yAxisValues[j]
	})
	yLabels := goplotter.NewLabels(yAxisValues, sp.YLabelColor)

	xAxis := goplotter.NewAxis(xLabels, 3, sp.XLabelColor)
	yAxis := goplotter.NewAxis(yLabels, 3, sp.XLabelColor)

	plot := goplotter.NewPlot(sp.Width, sp.Height, xAxis, yAxis)
	plot.BackgroundColor = sp.BackgroundColor
	plot.Title = sp.Title
	plot.TitleColor = sp.TitleColor

	plot.Draw()

	for i := 0; i < len(sp.XValues); i++ {
		xPosition := plot.AxisX.Labels[i].Position
		var yPosition int
		fmt.Println(sp.YValues[i])
		for _, label := range plot.AxisY.Labels {
			if label.Name == sp.YValues[i] {
				yPosition = label.Position
			}
		}
		point := goplotter.NewPoint(xPosition, int(yPosition), sp.PointShape)
		plot.AddPoint(point)

	}

	return plot.Draw()
}

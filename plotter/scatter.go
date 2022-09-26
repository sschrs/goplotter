package plotter

import (
	"fmt"
	"image/color"
)

type ScatterPlotter struct {
	XValues, YValues                                     []string
	Title, XLabel, YLabel                                string
	BackroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	AddLegend                                            bool
}

func NewScatterPlot(x []interface{}, y []float64) *ScatterPlotter {
	var xLabels []string
	var yLabels []string

	for _, xValue := range x {
		xLabels = append(xLabels, fmt.Sprintf("%v", xValue))
	}

	for _, yValue := range y {
		yLabels = append(yLabels, fmt.Sprintf("%f", yValue))
	}

	return &ScatterPlotter{
		XValues: xLabels,
		YValues: yLabels,
	}
}

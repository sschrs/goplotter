package plotter

import (
	"github.com/sschrs/goplotter"
	"image/color"
)

type ScatterPlotter struct {
	XValues, YValues                                      []float64
	Title, XLabel, YLabel                                 string
	BackgroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	AddLegend                                             bool
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
		AddLegend:       false,
		Width:           1280,
		Height:          720,
		PointShape:      goplotter.NewCircle(5, color.RGBA{R: 255, A: 255}),
	}
}

func (sp *ScatterPlotter) Plot() {
	// Set XAxis Range
	if sp.XRange.Step == 0 {
		minX, maxX := minMax(sp.XValues)
		step := (maxX - minX) / 10
		sp.XRange.From = minX
		sp.XRange.To = maxX
		sp.XRange.Step = step
	}

	// Set YAxis Range
	if sp.YRange.Step == 0 {
		minY, maxY := minMax(sp.YValues)
		step := (maxY - minY) / 10
		sp.YRange.From = minY
		sp.YRange.To = maxY
		sp.YRange.Step = step
	}

}

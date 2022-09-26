package plotter

import "image/color"

type ScatterPlotter struct {
	XValues, YValues                                     []string
	Title, XLabel, YLabel                                string
	BackroundColor, TitleColor, XLabelColor, YLabelColor color.Color
	AddLegend                                            bool
}

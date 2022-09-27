package plotter

type Range struct {
	From, To, Step float64
}

func GenerateRange(from, to, step float64) *Range {
	return &Range{
		From: from,
		To:   to,
		Step: step,
	}
}

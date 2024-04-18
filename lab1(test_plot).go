package main

import (
	"comp_math/lab1"
	"math"
)

func SomeFunction(x float64) float64 {
	return math.Sqrt(math.Cos(x)) - math.Pow(x, 2)
}

func main() {
	plot := lab1.NewPlotter()
	startInterval := -1.5
	endInterval := 1.5
	accuracy := 0.01

	x := startInterval
	for x < endInterval {
		pair := []float64{x, SomeFunction(x)}
		plot.CollectPoints(pair)
		x += accuracy
	}

	plot.DrawByPoints()
	plot.ExportToPng()

	plot.SetOutputFilename("output2.png")
	plot.ImportFromTxt()
	plot.DrawByPoints()
	plot.ExportToPng()
}

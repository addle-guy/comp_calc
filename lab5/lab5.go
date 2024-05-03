package main

import (
	"comp_math/lab1"
	"fmt"
	"math"
)

const (
	startOfInterval = -1.0
	endOfInterval   = -0.1
	N               = 2
	accuracy        = 0.0001
)

func ExecFunc(x float64) float64 {
	return math.Sqrt(math.Cos(x)) - math.Pow(x, 2)
}

func main() {
	plot := lab1.NewPlotter()
	// Просто распечатаем граф
	x := startOfInterval
	for float64(x) <= endOfInterval {
		pair := []float64{float64(x), ExecFunc(x)}
		plot.CollectPoints(pair)
		x += accuracy
	}

	// Определем равномерные узлы
	var nodes []float64
	nodes = append(nodes, -0.6)
	nodes = append(nodes, 0.2)
	/* step := math.Abs(startOfInterval-endOfInterval) / N
	for len(nodes) < N {
		if len(nodes) == 0 {
			nodes = append(nodes, startOfInterval+step)
		} else {
			nodes = append(nodes, nodes[len(nodes)-1]+step)
		}
	}
	*/
	var rootsOfNodes []float64
	for _, node := range nodes {
		rootsOfNodes = append(rootsOfNodes, ExecFunc(float64(node)))
	}

	simpleInterpol := GetSimplenterPolation(nodes, rootsOfNodes)
	x = startOfInterval
	k := 0
	for x <= endOfInterval {
		pair := []float64{x, simpleInterpol[k]}
		plot.CollectAdditionalPoints(pair)
		x += accuracy
		k += 1
	}
	plot.DrawByPoints()
	plot.SetOutputFilename("lab5.png")
	plot.ExportToPng()
}

func GetSimplenterPolation(nodes []float64, rootsOfNodes []float64) []float64 {
	x := startOfInterval
	interpolationResult := []float64{}
	for x < endOfInterval {
		Pn := 0.0
		for i := 0; i < N; i++ {
			Yk := rootsOfNodes[i]
			UpperMember := 1.0
			for j := 0; j < N; j++ {
				if j == i {
					continue
				}
				UpperMember *= (x - nodes[j])
			}
			LowerMember := 1.0
			for j := 0; j < N; j++ {
				if j == i {
					continue
				}
				LowerMember *= (nodes[i] - nodes[j])
			}
			Pn += Yk * UpperMember / LowerMember
		}
		fmt.Println(Pn)
		interpolationResult = append(interpolationResult, Pn)
		x += accuracy
	}
	fmt.Println(interpolationResult)
	return interpolationResult
}

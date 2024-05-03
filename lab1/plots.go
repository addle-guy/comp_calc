package lab1

import (
	"bufio"
	"fmt"
	"github.com/StephaneBunel/bresenham"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

type Plotter struct {
	backgroundColor         color.Color
	axisColor               color.Color
	graphsColor             color.Color
	additionalColor         color.Color
	anotherAdditionalColor  color.Color
	inputFile               *os.File
	outputFile              *os.File
	outputFilename          string
	inputFilename           string
	size                    int
	scaleX                  int
	scaleY                  int
	notchSize               int
	points                  []*Point
	AdditionalPoints        []*Point
	anotherAdditionalPoints []*Point
	chart                   *image.RGBA
	middle                  *Point
}

func NewPlotter() *Plotter {
	return &Plotter{
		backgroundColor:        color.White,
		axisColor:              color.Black,
		graphsColor:            color.RGBA{R: 255, A: 255},
		additionalColor:        color.RGBA{B: 255, A: 255},
		anotherAdditionalColor: color.RGBA{B: 255, A: 255},
		outputFilename:         "output.png",
		inputFilename:          "input.txt",
		size:                   500,
		scaleX:                 10,
		scaleY:                 10,
		notchSize:              3,
	}
}

func (p *Plotter) SetBackgroundColor(color color.Color) {
	p.backgroundColor = color
}

func (p *Plotter) SetAxisColor(color color.Color) {
	p.axisColor = color
}

func (p *Plotter) SetInputFile(file *os.File) {
	p.inputFile = file
}

func (p *Plotter) SetOutputFile(file *os.File) {
	p.outputFile = file
}

func (p *Plotter) SetInputFilename(filename string) {
	p.inputFilename = filename
}

func (p *Plotter) SetOutputFilename(filename string) {
	p.outputFilename = filename
}

func (p *Plotter) SetSize(size int) {
	p.size = size
}

func (p *Plotter) SetScaleX(scaleX int) {
	p.scaleX = scaleX
}

func (p *Plotter) SetScaleY(scaleY int) {
	p.scaleY = scaleY
}
func (p *Plotter) SetNotchSize(notchSize int) {
	p.notchSize = notchSize
}

func (p *Plotter) CollectPoints(data []float64) {
	p.points = append(p.points, &Point{x: int(data[0] * float64(p.size/p.scaleX)), y: int(data[1] * float64(p.size/p.scaleY))})
}

func (p *Plotter) CollectAdditionalPoints(data []float64) {
	p.AdditionalPoints = append(p.AdditionalPoints, &Point{x: int(data[0] * float64(p.size/p.scaleX)), y: int(data[1] * float64(p.size/p.scaleY))})
}

func (p *Plotter) CollectAnotherAdditionalPoints(data []float64) {
	p.points = append(p.anotherAdditionalPoints, &Point{x: int(data[0] * float64(p.size/p.scaleX)), y: int(data[1] * float64(p.size/p.scaleY))})
}

func (p *Plotter) AcceptPoints(points []*Point) {
	p.points = points
}

func (p *Plotter) ClearPoints() {
	p.points = []*Point{}
	p.AdditionalPoints = []*Point{}
	p.anotherAdditionalPoints = []*Point{}
}

func (p *Plotter) createChart() {
	p.chart = image.NewRGBA(image.Rect(0, 0, p.size, p.size))
	p.middle = NewPoint(p.size/2, p.size/2)
	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			p.chart.Set(i, j, color.White)
		}
	}
	bresenham.DrawLine(p.chart, 0, p.size/2, p.size, p.size/2, color.Black)
	bresenham.DrawLine(p.chart, p.size/2, 0, p.size/2, p.size, color.Black)
	for pen := p.middle.x; pen <= p.size; pen = pen + (p.size / p.scaleX) {
		bresenham.DrawLine(p.chart, pen, p.middle.y, pen, p.middle.y-p.notchSize, color.Black)
	}
	for pen := p.middle.x; pen >= 0; pen = pen - (p.size / p.scaleX) {
		bresenham.DrawLine(p.chart, pen, p.middle.y, pen, p.middle.y-p.notchSize, color.Black)
	}
	for pen := p.middle.y; pen <= p.size; pen = pen + (p.size / p.scaleY) {
		bresenham.DrawLine(p.chart, p.middle.x, pen, p.middle.x+p.notchSize, pen, color.Black)
	}
	for pen := p.middle.y; pen >= 0; pen = pen - (p.size / p.scaleY) {
		bresenham.DrawLine(p.chart, p.middle.x, pen, p.middle.x+p.notchSize, pen, color.Black)
	}
}

func (p *Plotter) DrawByPoints() {
	if len(p.points) <= 1 {
		fmt.Println("No points available")
		return
	}
	p.createChart()
	for i := 1; i < len(p.points); i++ {
		bresenham.DrawLine(p.chart,
			p.middle.x+p.points[i-1].x,
			p.middle.y-p.points[i-1].y,
			p.middle.x+p.points[i].x,
			p.middle.y-p.points[i].y,
			p.graphsColor)
	}
	for i := 1; i < len(p.AdditionalPoints); i++ {
		bresenham.DrawLine(p.chart,
			p.middle.x+p.AdditionalPoints[i-1].x,
			p.middle.y-p.AdditionalPoints[i-1].y,
			p.middle.x+p.AdditionalPoints[i].x,
			p.middle.y-p.AdditionalPoints[i].y,
			p.additionalColor)
	}

}

func (p *Plotter) ExportToPng() {
	var err error
	file, err := os.Create(p.outputFilename)
	if err != nil {
		defer fmt.Println("ne sozdal")
		panic(err)
	}
	err = png.Encode(file, p.chart)
	if err != nil {
		defer fmt.Println("tut zalupa")
		panic(err)
	}

}

func (p *Plotter) ImportFromTxt() {
	p.ClearPoints()
	var err error
	p.inputFile, err = os.Open(p.inputFilename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(p.inputFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		}
		fmt.Println(line)
		strs := strings.Split(strings.TrimSpace(line), " ")
		fmt.Println("1:", strs[0], "2:", strs[1])
		var values []float64
		for _, str := range strs {
			elem, err := strconv.ParseFloat(str, 64)
			if err != nil {
				panic(err)
			}
			values = append(values, elem)
		}
		fmt.Println(values)
		fmt.Println(int(values[0] * float64(p.size/p.scaleX)))
		fmt.Println(int(values[1] * float64(p.size/p.scaleY)))
		p.points = append(p.points, &Point{x: int(values[0] * float64(p.size/p.scaleX)), y: int(values[1] * float64(p.size/p.scaleY))})

	}
}

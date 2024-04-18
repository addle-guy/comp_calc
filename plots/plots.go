package main

import (
	"github.com/StephaneBunel/bresenham"
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	scale      = 10
	size       = 500
	notchSize  = 3
	whiteIndex = 0
	greyIndex  = 1
	blackIndex = 2
	redIndex   = 3
	blueIndex  = 4
	greenIndex = 5
)

type Point struct {
	x int
	y int
}

func main() {
	chart := image.NewRGBA(image.Rect(0, 0, size, size))
	// Заполняем холст
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			chart.Set(i, j, color.White)
		}
	}
	// Рисуем оси
	bresenham.DrawLine(chart, 0, size/2, size, size/2, color.Black)
	bresenham.DrawLine(chart, size/2, 0, size/2, size, color.Black)
	// Рисуем риски
	middle := Point{size / 2, size / 2}
	for pen := middle.x; pen <= size; pen = pen + (size / scale) {
		bresenham.DrawLine(chart, pen, middle.y, pen, middle.y-notchSize, color.Black)
	}
	for pen := middle.x; pen >= 0; pen = pen - (size / scale) {
		bresenham.DrawLine(chart, pen, middle.y, pen, middle.y-notchSize, color.Black)
	}
	for pen := middle.y; pen <= size; pen = pen + (size / scale) {
		bresenham.DrawLine(chart, middle.x, pen, middle.x+notchSize, pen, color.Black)
	}
	for pen := middle.y; pen >= 0; pen = pen - (size / scale) {
		bresenham.DrawLine(chart, middle.x, pen, middle.x+notchSize, pen, color.Black)
	}

	var x float64
	x = -10.0
	accuracy := 0.1
	current := Point{int(x * size / scale), int(x * x)}
	for x < 5 {
		next := Point{int((x + accuracy) * size / scale), int((x + accuracy) * (x + accuracy) * size / scale)}
		bresenham.DrawLine(chart,
			middle.x-current.x,
			middle.y-current.y,
			middle.x-next.x,
			middle.y-next.y,
			color.RGBA{255, 0, 0, 255})
		current = next
		x += accuracy
	}

	toimg, _ := os.Create("example1.png")
	defer toimg.Close()
	png.Encode(toimg, chart)
}

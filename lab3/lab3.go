// Метод золого сечения
// Функция: f(x) = sqrt(cos(x)) - x^2

package main

import (
	"fmt"
	"math"
)

// Начальные условия
const (
	accuracy        = 0.0000001
	Fi              = 1.6180339887
	startOfInterval = -2.0
	endOfInterval   = 2.0
)

// Функция
func ExecFunc(x float64) float64 {
	return math.Sqrt(math.Cos(x)) - math.Pow(x, 2)
}

func main() {
	a := startOfInterval
	b := endOfInterval
	counter := 0
	// Вычисление методом золотого сечения с заданной точностью
	for math.Abs(b-a) > accuracy {
		x1 := b - (b-a)/Fi
		x2 := a + (b-a)/Fi
		fx1 := ExecFunc(x1)
		fx2 := ExecFunc(x2)

		if fx1 <= fx2 {
			a = x1
		} else {
			b = x2
		}

		counter = counter + 1
	}
	fmt.Printf("Экстремум функции: %.10f, число иттераций: %v, точность: %v",
		(a+b)/2,
		counter,
		accuracy)
}

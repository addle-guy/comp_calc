// Метод Ньютона разностный аналог (метод секущих)
// Функция: f(x) = sqrt(cos(x)) - x^2

package main

import (
	"fmt"
	"math"
)

// Начальные условия
const (
	initialApproxOfFirst  = 0
	initialApproxOfSecond = 1
	accuracy              = 0.0001
)

// Функция
func ExecFunc(x float64) float64 {
	return math.Sqrt(math.Cos(x)) - math.Pow(x, 2)
}

// Поиск корня
func FindRoot(approx float64) (float64, int) {
	var prevRoot, currentRoot, nextRoot float64

	// Задаём начальное приближение
	prevRoot = approx
	currentRoot = approx - accuracy

	// Поиск Xk+1
	nextRoot = currentRoot - ExecFunc(currentRoot)*(currentRoot-prevRoot)/
		(ExecFunc(currentRoot+(currentRoot-prevRoot))-ExecFunc(prevRoot))
	counter := 1
	// Циклический поиск корня по заданной точности
	for math.Abs(nextRoot-currentRoot) > accuracy {
		prevRoot = currentRoot
		currentRoot = nextRoot
		nextRoot = currentRoot - ExecFunc(currentRoot)*(currentRoot-prevRoot)/
			(ExecFunc(currentRoot+(currentRoot-prevRoot))-ExecFunc(prevRoot))
		counter = counter + 1
	}
	return nextRoot, counter
}

func main() {
	// Поиск первого корня
	firstRoot, firstCounter := FindRoot(initialApproxOfFirst)
	fmt.Printf("Первый корень: %v, число иттераций: %v, точность: %v\n",
		firstRoot,
		firstCounter,
		accuracy)
	// Поиск второго корня
	secondRoot, secondCounter := FindRoot(initialApproxOfSecond)
	fmt.Printf("Ворой корень: %v, число иттераций: %v, точность: %v\n",
		secondRoot,
		secondCounter,
		accuracy)
}

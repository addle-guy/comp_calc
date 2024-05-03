// Решение СЛАУ методом Гаусса с постолбцовым выбором ведущего элемента
// Матрица и вектор в input.txt

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	filename = "./input.txt"
)

func main() {
	a, b, _, err := GetMatrix("./lab4/input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

// Чтение из файла с помощью bufio.NewReader
func GetMatrix(filename string) ([][]float64, []float64, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, 0, err
	}
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, 0, err
	}
	size, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, nil, 0, err
	}

	var a [][]float64
	var b []float64
	for i := 0; i < size; i++ {
		var buff []float64
		line, err = reader.ReadString('\n')
		chunks := strings.Split(strings.TrimSpace(line), " ")
		for n, chunk := range chunks {
			elem, err := strconv.ParseFloat(chunk, 64)
			if err != nil {
				return nil, nil, 0, err
			}
			if n == size-1 {
				b = append(b, elem)
			} else {
				buff = append(buff, elem)
			}
		}
		a = append(a, buff)
	}
	return a, b, size, nil
}

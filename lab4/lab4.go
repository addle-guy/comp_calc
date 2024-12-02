// Решение СЛАУ методом Гаусса с постолбцовым выбором ведущего элемента
// Матрица и вектор в input.txt
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "./lab4/input.txt"
const labda = 1e-9

func main() {
	mtxSize, mtx, vector, err := readInput(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:")
		log.Fatal(err)
	}
	fmt.Println("Матрица прочитана из файла " + filename)
	fmt.Println("Размерность матрицы:")
	fmt.Printf("%d\n", mtxSize)
	fmt.Println("Исходная расширенная матрица:")
	printMtx(mtx, vector)

	// Преобразование матрицы
	passRows := make(map[int]bool) // для пропуска ведущих строк
	for step := 1; step <= mtxSize; step++ {
		fmt.Printf("Преобразование: шаг %d\n", step)
		// Выбераем ведущую строку, ищем максимальный по модулю элемент матрицы
		max_i, max_j := 0, 0
		for i := 0; i < mtxSize; i++ {
			if passRows[i] {
				continue
			}
			for j := 0; j < mtxSize; j++ {
				if math.Abs(mtx[i][j]) > math.Abs(mtx[max_i][max_j]) {
					max_i, max_j = i, j
				}
			}
		}
		// Зафиксируем ведущую строку
		passRows[max_i] = true
		// Делим все элементы строки с максимальным элементом на максимальный элемент
		c := mtx[max_i][max_j]
		vector[max_i] /= c
		for j := 0; j < mtxSize; j++ {
			mtx[max_i][j] /= c
		}
		// Преобразовываем остальные строки
		for i := 0; i < mtxSize; i++ {
			if passRows[i] {
				continue
			}
			c := mtx[i][max_j]
			vector[i] -= vector[max_i] * c
			for j := 0; j < mtxSize; j++ {
				mtx[i][j] -= mtx[max_i][j] * c
			}
		}

		printMtx(mtx, vector)
	}
	// Выстроим строки
	for j := 0; j < mtxSize; j++ {
		for i := j; i < mtxSize; i++ {
			if (mtx[i][j] > 1.0-labda) && (mtx[i][j] < 1.0+labda) {
				mtx[i], mtx[j] = mtx[j], mtx[i]
				vector[i], vector[j] = vector[j], vector[i]
			}
		}
	}

	fmt.Println("Результат преобразования:")
	printMtx(mtx, vector)

	// Выставим порядок поиска
	for i := 0; i < mtxSize; i++ {
		for k := i + 1; k < mtxSize; k++ {
			countA := 0
			countB := 0
			for j := 0; j < mtxSize; j++ {
				if math.Abs(mtx[i][j]) < labda {
					countA++
				}
				if math.Abs(mtx[k][j]) < labda {
					countB++
				}
			}
			if countA < countB {
				mtx[i], mtx[k] = mtx[k], mtx[i]
				vector[i], vector[k] = vector[k], vector[i]
			}
		}
	}
	fmt.Println("Для тестирования (порядок поиска):")
	printMtx(mtx, vector)

	// Поиск решения
	solves := make(map[int]float64)

	// Ищем корни
	for len(solves) < mtxSize {
		for i := 0; i < mtxSize; i++ {
			tg := 0
			for j := 0; j < mtxSize; j++ {
				if (math.Abs(mtx[i][j]) > 1.0-labda) && (math.Abs(mtx[i][j]) < 1.0+labda) {
					tg = j
					solves[tg] = vector[i]
				}
			}
			for j := 0; j < mtxSize; j++ {
				if math.Abs(mtx[i][j]) < labda {
					continue
				}
				if j != tg {
					solves[tg] = solves[tg] - mtx[i][j]*solves[j]
				}
			}
		}
	}

	fmt.Println("Решения:")
	var order []int
	for k, _ := range solves {
		order = append(order, k)
	}
	sort.Ints(order)

	for n := range order {
		fmt.Printf("x%d = %8.4f\n", n+1, solves[order[n]])
	}
	return
}

// Читает файл input.txt
func readInput(filename string) (int, [][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()

	var size int64
	var mtx [][]float64
	var vector []float64

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	size, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 32)
	if err != nil {
		return 0, nil, nil, err
	}
	for scanner.Scan() {
		var members []float64
		line := strings.TrimSpace(scanner.Text())
		strMembers := strings.Split(line, " ")
		for n, strMember := range strMembers {
			member, err := strconv.ParseFloat(strMember, 64)
			if err != nil {
				return 0, nil, nil, err
			}
			if n == int(size) {
				vector = append(vector, member)
				break
			}
			members = append(members, member)
		}
		mtx = append(mtx, members)
	}

	return int(size), mtx, vector, nil
}

// Печать матрицы на экран
func printMtx(mtx [][]float64, vector []float64) {
	cntVector := 0
	for _, row := range mtx {
		for _, member := range row {
			fmt.Printf("%8.4f", member)
		}
		fmt.Printf("  | %4.4f\n", vector[cntVector])
		cntVector++
	}
	fmt.Println()
}

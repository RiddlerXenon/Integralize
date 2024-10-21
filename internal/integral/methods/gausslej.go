package methods

import (
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Вычисление моментов весовой функции на интервале [-1, 1]
func computeMoments(n int) []float64 {
	moments := make([]float64, n)
	for j := 0; j < n; j++ {
		moments[j] = 2 / float64(j+1) // Моменты для p(x) = 1 на [-1, 1]
	}
	return moments
}

// Решение системы уравнений для коэффициентов многочлена
func solveSystem(n int, moments []float64) []float64 {
	A := mat.NewDense(n, n, nil)
	b := mat.NewVecDense(n, moments)

	// Заполнение матрицы системы
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.Set(i, j, math.Pow(float64(j+1), float64(i)))
		}
	}

	// Решение системы A * x = b
	var weights mat.VecDense
	err := weights.SolveVec(A, b)
	if err != nil {
		log.Fatalf("Ошибка решения системы: %v", err)
	}

	// Преобразование в массив весов
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = weights.AtVec(i)
	}

	return result
}

// Вспомогательная функция для вычисления полинома Лежандра
func legendrePolynomial(n int, x float64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return ((2*float64(n)-1)*x*legendrePolynomial(n-1, x) - (float64(n)-1)*legendrePolynomial(n-2, x)) / float64(n)
}

// Поиск корней полинома Лежандра (узлы) для метода Гаусса
func findLegendreRoots(n int) []float64 {
	tolerance := 1e-14
	roots := make([]float64, n)

	// Поиск корней многочлена Лежандра
	for i := 0; i < n; i++ {
		x := math.Cos(math.Pi * (float64(i) + 0.75) / (float64(n) + 0.5)) // Начальное приближение
		for {
			p := legendrePolynomial(n, x)
			dp := (float64(n) * (legendrePolynomial(n-1, x) - x*p)) / (1 - x*x) // Производная
			delta := p / dp
			x -= delta
			if math.Abs(delta) < tolerance {
				break
			}
		}
		roots[i] = x
	}

	return roots
}

// Основная функция для интегрирования методом Гаусса
func GaussQuadrature(a, b float64, n int, expr string) (float64, error) {
	// Вычисляем узлы и веса
	nodes := findLegendreRoots(n)
	weights := solveSystem(n, computeMoments(n))

	mid := (a + b) / 2.0
	halfLength := (b - a) / 2.0

	sum := 0.0
	for i := 0; i < n; i++ {
		x := mid + halfLength*nodes[i]
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += weights[i] * fx
	}

	return halfLength * sum, nil
}

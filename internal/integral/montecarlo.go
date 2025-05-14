package integral

import (
	"errors"
	"math"
	"math/rand"
	"sync"
)

// MonteCarlo вычисляет интеграл функции f от a до b методом Монте-Карло
// с погрешностью не более 10%
func MonteCarlo(a, b float64, initialN float64, f func(vars map[string]float64) float64) (float64, error) {
	if a >= b {
		return 0, errors.New("неверные границы интегрирования: a должно быть меньше b")
	}

	if initialN <= 0 {
		return 0, errors.New("количество точек должно быть положительным")
	}

	// Первое приближение
	result1, err := monteCarloCalculation(a, b, initialN, f)
	if err != nil {
		return 0, err
	}

	// Второе приближение с удвоенным числом точек
	n := initialN * 4
	result2, err := monteCarloCalculation(a, b, n, f)
	if err != nil {
		return 0, err
	}

	// Проверка погрешности
	for math.Abs(result2-result1)/math.Abs(result2) > 0.1 {
		result1 = result2
		n *= 4
		// Ограничение максимального числа итераций
		if n > 100000000 {
			return result2, errors.New("достигнуто максимальное количество итераций, но требуемая точность не достигнута")
		}
		result2, err = monteCarloCalculation(a, b, n, f)
		if err != nil {
			return 0, err
		}
	}

	return result2, nil
}

// Вспомогательная функция для параллельного вычисления интеграла
func monteCarloCalculation(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	// Используем несколько горутин для параллельных вычислений
	numGoroutines := 4
	pointsPerGoroutine := int(n) / numGoroutines

	var wg sync.WaitGroup
	var mu sync.Mutex
	sum := 0.0

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(seed int64) {
			defer wg.Done()

			// Создаем генератор случайных чисел с уникальным seed для каждой горутины
			localRNG := rand.New(rand.NewSource(seed))
			localSum := 0.0

			for j := 0; j < pointsPerGoroutine; j++ {
				x := map[string]float64{
					"x": a + localRNG.Float64()*(b-a),
				}
				fx := f(x)
				localSum += fx
			}

			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(rand.Int63() + int64(i))
	}

	wg.Wait()

	// Обработка оставшихся точек (если n не делится нацело на количество горутин)
	remainingPoints := int(n) % numGoroutines
	if remainingPoints > 0 {
		localRNG := rand.New(rand.NewSource(rand.Int63()))
		for i := 0; i < remainingPoints; i++ {
			x := map[string]float64{
				"x": a + localRNG.Float64()*(b-a),
			}
			sum += f(x)
		}
	}

	return (b - a) * sum / float64(n), nil
}

// MonteCarloAdaptive вычисляет интеграл с адаптивным разбиением интервала
func MonteCarloAdaptive(a, b float64, initialN int, f func(vars map[string]float64) float64) (float64, error) {
	if a >= b {
		return 0, errors.New("неверные границы интегрирования: a должно быть меньше b")
	}

	if initialN <= 0 {
		return 0, errors.New("количество точек должно быть положительным")
	}

	// Разделим интервал на подинтервалы
	numSubintervals := 10
	subintervalLength := (b - a) / float64(numSubintervals)

	var wg sync.WaitGroup
	results := make([]float64, numSubintervals)
	errors := make([]error, numSubintervals)

	// Вычислим интеграл на каждом подинтервале параллельно
	for i := 0; i < numSubintervals; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			subA := a + float64(idx)*subintervalLength
			subB := subA + subintervalLength
			pointsPerSubinterval := initialN / numSubintervals
			result, err := MonteCarlo(subA, subB, float64(pointsPerSubinterval), f)
			results[idx] = result
			errors[idx] = err
		}(i)
	}

	wg.Wait()

	// Проверим ошибки
	for _, err := range errors {
		if err != nil {
			return 0, err
		}
	}

	// Суммируем результаты подинтегралов
	totalResult := 0.0
	for _, result := range results {
		totalResult += result
	}

	return totalResult, nil
}

package integral

// SimpsonMethod использует метод Симпсона для численного интегрирования
func Simpson(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	// Убеждаемся, что n четное
	if int(n)%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)

	// Начальное и конечное значение функции
	fa := f(map[string]float64{
		"x": a,
	})
	fb := f(map[string]float64{
		"x": b,
	})

	// Вычисление начальной суммы
	sum := fa + fb

	// Считаем для нечётных индексов
	for i := 1.0; i < n; i += 2.0 {
		fi := f(map[string]float64{
			"x": a + i*h,
		})
		sum += 4 * fi
	}

	// Считаем для чётных индексов
	for i := 2.0; i < n; i += 2.0 {
		fi := f(map[string]float64{
			"x": a + i*h,
		})
		sum += 2 * fi
	}

	// Возвращаем результат интегрирования
	return sum * h / 3.0, nil
}

package methods

// SimpsonMethod использует метод Симпсона для численного интегрирования
func SimpsonMethod(a, b, n float64, expr string) (float64, error) {
	// Убеждаемся, что n четное

	if int(n)%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)

	// Начальное и конечное значение функции
	fa, err := f(a, expr)
	if err != nil {
		return 0, err
	}
	fb, err := f(b, expr)
	if err != nil {
		return 0, err
	}

	// Вычисление начальной суммы
	sum := fa + fb

	// Считаем для нечётных индексов
	for i := 1.0; i < n; i += 2.0 {
		fi, err := f(a+i*h, expr)
		if err != nil {
			return 0, err
		}
		sum += 4 * fi
	}

	// Считаем для чётных индексов
	for i := 2.0; i < n; i += 2.0 {
		fi, err := f(a+i*h, expr)
		if err != nil {
			return 0, err
		}
		sum += 2 * fi
	}

	// Возвращаем результат интегрирования
	return sum * h / 3.0, nil
}

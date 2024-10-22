package methods

// TrapezoidMethod использует метод трапеций для численного интегрирования
func TrapezoidMethod(a, b, n float64, expr string) (float64, error) {
	h := (b - a) / n

	// Начальные значения функции
	fa, err := f(a, expr)
	if err != nil {
		return 0, err
	}
	fb, err := f(b, expr)
	if err != nil {
		return 0, err
	}

	// Начальная сумма
	sum := (fa + fb) / 2.0

	// Основной цикл по интервалам
	for i := 1.0; i < n; i++ {
		x := a + i*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	// Возвращаем результат интегрирования
	return sum * h, nil
}

package integral

// TrapezoidMethod использует метод трапеций для численного интегрирования
func Trapezoidal(a, b, n float64, f func(x float64) float64) (float64, error) {
	h := (b - a) / n

	// Начальные значения функции
	fa := f(a)
	fb := f(b)

	// Начальная сумма
	sum := (fa + fb) / 2.0

	// Основной цикл по интервалам
	for i := 1.0; i < n; i++ {
		x := a + i*h
		fx := f(x)
		sum += fx
	}

	// Возвращаем результат интегрирования
	return sum * h, nil
}

package integral

// TrapezoidMethod использует метод трапеций для численного интегрирования
func Trapezoidal(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	h := (b - a) / n

	// Начальные значения функции
	fa := f(map[string]float64{
		"x": a,
	})
	fb := f(map[string]float64{
		"x": b,
	})

	// Начальная сумма
	sum := (fa + fb) / 2.0

	// Основной цикл по интервалам
	for i := 1.0; i < n; i++ {
		x := map[string]float64{
			"x": a + i*h,
		}
		fx := f(x)
		sum += fx
	}

	// Возвращаем результат интегрирования
	return sum * h, nil
}

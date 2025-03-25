package integral

// Метод левых прямоугольников
func LeftRectangleMethod(a, b, n float64, f func(x float64) float64) (float64, error) {
	// Шаг разбиения
	h := (b - a) / n
	sum := 0.0

	// Суммирование значений функции в левых точках каждого отрезка
	for i := 0; i < int(n); i++ {
		x := a + float64(i)*h
		sum += f(x)
	}

	return sum * h, nil
}

// Метод правых прямоугольников
func RightRectangleMethod(a, b, n float64, f func(x float64) float64) (float64, error) {
	// Шаг разбиения
	h := (b - a) / n
	sum := 0.0

	// Суммирование значений функции в правых точках каждого отрезка
	for i := 1; i <= int(n); i++ {
		x := a + float64(i)*h
		sum += f(x)
	}

	return sum * h, nil
}

func MidpointRectangleMethod(a, b, n float64, f func(x float64) float64) (float64, error) {
	h := (b - a) / n
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + (i+0.5)*h
		sum += f(x)
	}

	return sum * h, nil
}

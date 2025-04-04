package integral

// Метод левых прямоугольников
func LeftRectangle(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	// Шаг разбиения
	h := (b - a) / n
	sum := 0.0

	// Суммирование значений функции в левых точках каждого отрезка
	for i := 0; i < int(n); i++ {
		x := map[string]float64{
			"x": a + float64(i)*h,
		}
		sum += f(x)
	}

	return sum * h, nil
}

// Метод правых прямоугольников
func RightRectangle(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	// Шаг разбиения
	h := (b - a) / n
	sum := 0.0

	// Суммирование значений функции в правых точках каждого отрезка
	for i := 1; i <= int(n); i++ {
		x := map[string]float64{
			"x": a + float64(i)*h,
		}

		sum += f(x)
	}

	return sum * h, nil
}

func MidpointRectangle(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	h := (b - a) / n
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := map[string]float64{
			"x": a + (i+0.5)*h,
		}
		sum += f(x)
	}

	return sum * h, nil
}

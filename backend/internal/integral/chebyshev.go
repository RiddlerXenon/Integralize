package integral

import (
	"errors"
	"math"
)

func Chebyshev(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}
	if a >= b {
		return 0, errors.New("a must be less than b")
	}

	sum := 0.0
	vars := make(map[string]float64)

	// Преобразование интервала [a, b] к [-1, 1]
	c1 := (b - a) / 2
	c2 := (a + b) / 2

	for k := 1.0; k <= n; k++ {
		// Узлы Чебышёва (корни полинома Чебышёва) на интервале [-1, 1]
		xk := math.Cos((2*k - 1) * math.Pi / (2 * n))

		// Преобразование узла к исходному интервалу [a, b]
		xTransformed := c1*xk + c2
		vars["x"] = xTransformed

		// Вычисление значения функции в узле
		fk := f(vars)

		// Суммирование с весами (все веса равны для метода Чебышёва)
		sum += fk
	}

	// Окончательное вычисление интеграла
	integral := (b - a) * sum / n
	return integral, nil
}

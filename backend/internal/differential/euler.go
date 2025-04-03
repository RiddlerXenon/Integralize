package differential

func Euler(y0, t0, tMax, h float64, f func(vars map[string]float64) float64) ([]float64, []float64) {
	//Разбиваем на шаги
	nSteps := int((tMax - t0) / h)

	/* Так как изначально диффур принимает вид как dy/dt,
	то логично использовать X = t, Y = y */
	t := make([]float64, nSteps+1)
	y := make([]float64, nSteps+1)

	t[0] = t0 //Начальное время
	y[0] = y0 //y(0) = Введенному значению

	/* Воспользуемся формулой
	y[i + 1] = y[i] - h * k * y[i]
	t[i + 1] = t[i] + h
	где k - коэффициент, y0 - начальное значение
	tMax - максимальное время, h - шаг */
	for i := 0; i < nSteps; i++ {
		t[i+1] = t[i] + h
		y[i+1] = y[i] + h*f(map[string]float64{
			"x": t[i],
			"y": y[i]})
	}

	return t, y
}

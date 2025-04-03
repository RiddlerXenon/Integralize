package differential

func RungeKutte(y0, t0, tMax, h float64, f func(vars map[string]float64) float64) ([]float64, []float64) {
	nSteps := int((tMax - t0) / h)
	t := make([]float64, nSteps+1)
	y := make([]float64, nSteps+1)

	t[0] = t0
	y[0] = y0

	for i := 0; i < nSteps; i++ {
		t[i+1] = t[i] + h
		k1 := h * f(map[string]float64{
			"x": t[i],
			"y": y[i]})
		k2 := h * f(map[string]float64{
			"x": t[i] + h/2,
			"y": y[i] + k1/2})
		k3 := h * f(map[string]float64{
			"x": t[i] + h/2,
			"y": y[i] + k2/2})
		k4 := h * f(map[string]float64{
			"x": t[i] + h,
			"y": y[i] + k3})

		y[i+1] = y[i] + (k1+2*k2+2*k3+k4)/6
	}

	return t, y
}

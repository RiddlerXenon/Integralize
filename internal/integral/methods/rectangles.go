package methods

func LeftRectangleMethod(a, b, n float64, expr string) (float64, error) {
	h := (b - a) / n
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + i*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

func RightRectangleMethod(a, b, n float64, expr string) (float64, error) {
	h := (b - a) / n
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + i*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

func MidpointRectangleMethod(a, b, n float64, expr string) (float64, error) {
	h := (b - a) / n
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + (i+0.5)*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

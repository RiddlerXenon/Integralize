package methods

func LeftRectangleMethod(a, b float64, n int, expr string) (float64, error) {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

func RightRectangleMethod(a, b float64, n int, expr string) (float64, error) {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 1; i <= n; i++ {
		x := a + float64(i)*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

func MidpointRectangleMethod(a, b float64, n int, expr string) (float64, error) {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*h
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return sum * h, nil
}

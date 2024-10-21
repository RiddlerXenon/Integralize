package methods

func LeftRectangleMethod(fp *FunctionParser, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		sum += fp.F(x)
	}

	return sum * h
}

func RightRectangleMethod(fp *FunctionParser, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 1; i <= n; i++ {
		x := a + float64(i)*h
		sum += fp.F(x)
	}

	return sum * h
}

func MidpointRectangleMethod(fp *FunctionParser, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*h
		sum += fp.F(x)
	}

	return sum * h
}

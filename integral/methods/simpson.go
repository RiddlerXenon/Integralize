package methods

func SimpsonMethod(fp *FunctionParser, a, b float64, n int) float64 {
	if n%2 != 0 {
		n++
	}
	h := (b - a) / float64(n)
	sum := fp.F(a) + fp.F(b)

	for i := 1; i < n; i += 2 {
		sum += 4 * fp.F(a+float64(i)*h)
	}

	for i := 2; i < n-1; i += 2 {
		sum += 2 * fp.F(a+float64(i)*h)
	}

	return sum * h / 3.0
}

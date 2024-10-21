package methods

func TrapezoidMethod(fp *FunctionParser, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := (fp.F(a) + fp.F(b)) / 2.0

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += fp.F(x)
	}

	return sum * h
}

package methods

import (
	"math/rand"
)

func MonteCarloMethod(fp *FunctionParser, a, b float64, n int) float64 {
	rng := rand.New(rand.NewSource(rand.Int63()))
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + rng.Float64()*(b-a)
		sum += fp.F(x)
	}

	return (b - a) * sum / float64(n)
}

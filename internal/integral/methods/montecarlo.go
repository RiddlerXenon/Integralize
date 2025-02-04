package methods

import (
	"math/rand"
)

func MonteCarloMethod(a, b, n float64, f func(x float64) float64) (float64, error) {
	rng := rand.New(rand.NewSource(rand.Int63()))
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + rng.Float64()*(b-a)
		fx := f(x)
		sum += fx
	}

	return (b - a) * sum / float64(n), nil
}

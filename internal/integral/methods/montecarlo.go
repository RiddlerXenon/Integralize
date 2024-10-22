package methods

import (
	"math/rand"
)

func MonteCarloMethod(a, b, n float64, expr string) (float64, error) {
	rng := rand.New(rand.NewSource(rand.Int63()))
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := a + rng.Float64()*(b-a)
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return (b - a) * sum / float64(n), nil
}

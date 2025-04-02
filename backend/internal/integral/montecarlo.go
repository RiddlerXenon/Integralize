package integral

import (
	"math/rand"
)

func MonteCarlo(a, b, n float64, f func(vars map[string]float64) float64) (float64, error) {
	rng := rand.New(rand.NewSource(rand.Int63()))
	sum := 0.0

	for i := 0.0; i < n; i++ {
		x := map[string]float64{
			"x": a + rng.Float64()*(b-a),
		}
		fx := f(x)
		sum += fx
	}

	return (b - a) * sum / float64(n), nil
}

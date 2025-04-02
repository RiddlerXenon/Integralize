package integral

import (
	"math"

	"go.uber.org/zap"
)

func chebyshevNodesWeights(n int) ([]float64, float64) {
	nodes := make([]float64, n)
	weights := math.Pi / float64(n)

	for i := 0; i < n; i++ {
		nodes[i] = math.Cos((2.0*float64(i+1) - 1.0) / (2.0 * float64(n)) * math.Pi)
	}

	return nodes, weights
}

func Chebyshev(a, b, n float64, f func(x float64) float64) (float64, error) {
	nodes, weight := chebyshevNodesWeights(int(n))

	mid := (a + b) / 2.0
	halfLength := (b - a) / 2.0

	sum := 0.0
	for i := 0; i < int(n); i++ {
		x := mid + halfLength*nodes[i]
		fx := f(x)
		sum += fx
	}

	zap.S().Infof("Chebyshev: %v", halfLength*weight*sum)

	return halfLength * weight * sum, nil
}

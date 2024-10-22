package methods

import "math"

func chebyshevNodesWeights(n int) ([]float64, float64) {
	nodes := make([]float64, n)
	weights := math.Pi / float64(n)

	for i := 0; i < n; i++ {
		nodes[i] = math.Cos((2.0*float64(i+1) - 1.0) / (2.0 * float64(n)) * math.Pi)
	}

	return nodes, weights
}

func ChebyshevQuadrature(a, b float64, n int, expr string) (float64, error) {
	nodes, weight := chebyshevNodesWeights(n)

	mid := (a + b) / 2.0
	halfLength := (b - a) / 2.0

	sum := 0.0
	for i := 0; i < n; i++ {
		x := mid + halfLength*nodes[i]
		fx, err := f(x, expr)
		if err != nil {
			return 0, err
		}
		sum += fx
	}

	return halfLength * weight * sum, nil
}

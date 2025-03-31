package differential

import (
	"errors"
	"math"
)

type Parameters struct {
	Alpha float64
	Beta  float64
	Delta float64
	Gamma float64
	Step  float64
	Steps int
	Prey  float64
	Pred  float64
}

func validateParams(p Parameters) error {
	if p.Alpha < 0 || p.Beta < 0 || p.Delta < 0 || p.Gamma < 0 || p.Step < 0 || p.Steps < 0 || p.Prey < 0 || p.Pred < 0 {
		return errors.New("parameters must be positive")
	}
	if p.Step <= 0 || p.Steps <= 0 {
		return errors.New("step and steps must be greater than 0")
	}
	if p.Prey < 0 || p.Pred < 0 {
		return errors.New("prey and pred must be greater than 0")
	}
	return nil
}

func EulerMethod(p Parameters) ([]float64, []float64, error) {
	if err := validateParams(p); err != nil {
		return nil, nil, err
	}
	prey := make([]float64, p.Steps)
	pred := make([]float64, p.Steps)
	prey[0] = p.Prey
	pred[0] = p.Pred
	for i := 1; i < p.Steps; i++ {
		dPrey := p.Alpha*prey[i-1] - p.Beta*prey[i-1]*pred[i-1]
		dPred := -p.Gamma*pred[i-1] + p.Delta*prey[i-1]*pred[i-1]

		prey[i] = prey[i-1] + p.Step*dPrey
		pred[i] = pred[i-1] + p.Step*dPred

		if prey[i] < 0 || math.IsNaN(prey[i]) || prey[i] < 0 || math.IsNaN(pred[i]) {
			return nil, nil, errors.New("invalid values. Population is negative")
		}
	}
	return prey, pred, nil
}

func RungeKuttaMethod(p Parameters) ([]float64, []float64, error) {
	if err := validateParams(p); err != nil {
		return nil, nil, err
	}

	prey := make([]float64, p.Steps)
	pred := make([]float64, p.Steps)

	prey[0] = p.Prey
	pred[0] = p.Pred

	for i := 1; i < p.Steps; i++ {
		h := p.Step

		k1Prey := h * (p.Alpha*prey[i-1] - p.Beta*prey[i-1]*pred[i-1])
		k1Pred := h * (-p.Gamma*pred[i-1] + p.Delta*prey[i-1]*pred[i-1])

		k2Prey := h * (p.Alpha*(prey[i-1]+k1Prey/2) - p.Beta*(prey[i-1]+k1Prey/2)*(pred[i-1]+k1Pred/2))
		k2Pred := h * (-p.Gamma*(pred[i-1]+k1Pred/2) + p.Delta*(prey[i-1]+k1Prey/2)*(pred[i-1]+k1Pred/2))

		k3Prey := h * (p.Alpha*(prey[i-1]+k2Prey/2) - p.Beta*(prey[i-1]+k2Prey/2)*(pred[i-1]+k2Pred/2))
		k3Pred := h * (-p.Gamma*(pred[i-1]+k2Pred/2) + p.Delta*(prey[i-1]+k2Prey/2)*(pred[i-1]+k2Pred/2))

		k4Prey := h * (p.Alpha*(prey[i-1]+k3Prey) - p.Beta*(prey[i-1]+k3Prey)*(pred[i-1]+k3Pred))
		k4Pred := h * (-p.Gamma*(pred[i-1]+k3Pred) + p.Delta*(prey[i-1]+k3Prey)*(pred[i-1]+k3Pred))

		prey[i] = prey[i-1] + (k1Prey+2*k2Prey+2*k3Prey+k4Prey)/6
		pred[i] = pred[i-1] + (k1Pred+2*k2Pred+2*k3Pred+k4Pred)/6

		if prey[i] < 0 || math.IsNaN(prey[i]) || prey[i] < 0 || math.IsNaN(pred[i]) {
			return nil, nil, errors.New("invalid values. Population is negative")
		}
	}

	return prey, pred, nil
}

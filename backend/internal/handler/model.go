package handler

import (
	//"github.com/RiddlerXenon/Integralize/internal/differential"
	"github.com/RiddlerXenon/Integralize/internal/differential"
	"github.com/RiddlerXenon/Integralize/internal/integral"
)

// Мапа для методов уравнений, ключ - название метода, значения - функция
var integralMethods = map[string]func(float64, float64, float64, func(map[string]float64) float64) (float64, error){
	"left rectangle":     integral.LeftRectangle,
	"right rectangle":    integral.RightRectangle,
	"midpoint rectangle": integral.MidpointRectangle,
	"trapezoidal":        integral.Trapezoidal,
	"simpson":            integral.Simpson,
	"monte carlo":        integral.MonteCarlo,
	"gauss lejandre":     integral.GaussLejandre,
	"chebyshev":          integral.Chebyshev,
}

var diffEquationsMethods = map[string]func(float64, float64, float64, float64, func(map[string]float64) float64) ([]float64, []float64){
	"euler":       differential.Euler,
	"runge-kutta": differential.RungeKutte,
}

var predatorVictim = map[string]func(differential.Parameters) ([]float64, []float64, error){
	"euler":       differential.EulerMethod,
	"runge-kutta": differential.RungeKuttaMethod,
}

// Структуры запросов для интегралов и дифференциальных уравнений
type integralRequest struct {
	EquationType string    `json:"equationType"`
	Expression   string    `json:"expression"`
	Args         []float64 `json:"args"`
}

type diffEquationsRequest struct {
	EquationType string    `json:"equationType"`
	Expression   string    `json:"expression"`
	Args         []float64 `json:"args"`
}

type predatorVictimRequest struct {
	EquationType string    `json:"equationType"`
	Args         []float64 `json:"args"`
	Step         float64   `json:"step"`
	Steps        int       `json:"steps"`
	PredVictim   []float64 `json:"pred_victim"`
}

type predatorVictimResponse struct {
	PredY []float64 `json:"predator"`
	PreyY []float64 `json:"prey"`
}

type integralResponse struct {
	Result float64 `json:"result"`
}

type diffEquationsResponse struct {
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

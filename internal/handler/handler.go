package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/diff_eq"
	integral "github.com/RiddlerXenon/Integralize/internal/integral/methods"
	"github.com/RiddlerXenon/Integralize/internal/parser"
)

// Типы данных для методов интегралов и дифференциальных уравнений
type IntegralMethod string
type DiffEquationsMethod string

const (
	// Методы интегралов
	LeftRectangle     IntegralMethod = "left rectangle"
	RightRectangle    IntegralMethod = "right rectangle"
	MidpointRectangle IntegralMethod = "midpoint rectangle"
	Trapezoidal       IntegralMethod = "trapezoidal"
	Simpson           IntegralMethod = "simpson"
	MonteCarlo        IntegralMethod = "monte carlo"
	GaussLejandre     IntegralMethod = "gauss lejandre"
	Chebychev         IntegralMethod = "chebyshev"

	// Методы дифференциальных уравнений
	Euler      DiffEquationsMethod = "euler"
	RungeKutta DiffEquationsMethod = "runge-kutta"
)

// Структуры запросов для интегралов и дифференциальных уравнений
type IntegralRequest struct {
	Method     IntegralMethod `json:"method"`
	Expression string         `json:"expression"`
	A          float64        `json:"a"`
	B          float64        `json:"b"`
	N          float64        `json:"n"`
}

type DiffEquationsRequest struct {
	Method     DiffEquationsMethod `json:"method"`
	Expression string              `json:"expression"`
}

// Функции для парсинга выражений
func parseIntegralExpression(expression string) (func(float64) float64, error) {
	return parser.ParseStrInt(expression)
}

func parseDiffEquationsExpression(expression string) (func(float64, float64) float64, error) {
	return parser.ParseStrDiffEq(expression)
}

// Хэндлеры для обработки запросов
func integralHandler(w http.ResponseWriter, r *http.Request) {
	var req IntegralRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Парсинг выражения
	expressionFunc, err := parseIntegralExpression(req.Expression)
	if err != nil {
		http.Error(w, "failed to parse expression", http.StatusBadRequest)
		return
	}

	// Обработка запроса с использованием распарсенного выражения и передачи в соответствующий метод
	switch req.Method {
	case LeftRectangle:
		integral.LeftRectangleMethod(req.A, req.B, req.N, expressionFunc)
	case RightRectangle:
		integral.RightRectangleMethod(req.A, req.B, req.N, expressionFunc)
	case MidpointRectangle:
		integral.MidpointRectangleMethod(req.A, req.B, req.N, expressionFunc)
	case Trapezoidal:
		integral.TrapezoidMethod(req.A, req.B, req.N, expressionFunc)
	case Simpson:
		integral.SimpsonMethod(req.A, req.B, req.N, expressionFunc)
	case MonteCarlo:
		integral.MonteCarloMethod(req.A, req.B, req.N, expressionFunc)
	case GaussLejandre:
		integral.GaussQuadrature(req.A, req.B, req.N, expressionFunc)
	case Chebychev:
		integral.ChebyshevQuadrature(req.A, req.B, req.N, expressionFunc)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Expression processed successfully")
}

func diffEquationsHandler(w http.ResponseWriter, r *http.Request) {
	var req DiffEquationsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Парсинг выражения
	expressionFunc, err := parseDiffEquationsExpression(req.Expression)
	if err != nil {
		http.Error(w, "failed to parse expression", http.StatusBadRequest)
		return
	}

	// Обработка запроса с использованием распарсенного выражения и передачи в соответствующий метод
	switch req.Method {
	case Euler:
		t, y := diff_eq.EulerMethod(expressionFunc, 1.0, 0.0, 2.0, 0.1) // пример вызова
		fmt.Fprintf(w, "Euler Method Result: t=%v, y=%v", t, y)
	case RungeKutta:
		t, y := diff_eq.RungeKutte(expressionFunc, 1.0, 0.0, 2.0, 0.1) // пример вызова
		fmt.Fprintf(w, "Runge-Kutta Method Result: t=%v, y=%v", t, y)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/api/integral", integralHandler)
	http.HandleFunc("/api/diff_equations", diffEquationsHandler)
	http.ListenAndServe(":8080", nil)
}

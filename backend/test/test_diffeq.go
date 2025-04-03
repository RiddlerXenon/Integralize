package main

import (
	"bufio"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Function representing dy/dx = f(x, y)
type Equation func(float64, float64) float64

// Function representing a second-order ODE d2y/dx2 = f(x, y, dy/dx)
type SecondOrderEquation func(float64, float64, float64) float64

// Функция для вычисления выражений через govaluate
type Evaluator struct {
	expr *govaluate.EvaluableExpression
}

func (e *Evaluator) Evaluate(params map[string]interface{}) float64 {
	result, err := e.expr.Evaluate(params)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return 0
	}
	return result.(float64)
}

// Функция для упрощения и преобразования уравнения в вычисляемую функцию
func simplifyEquation(equationStr string) (Equation, error) {
	equationStr = strings.ToLower(strings.ReplaceAll(equationStr, " ", ""))
	if !strings.HasPrefix(equationStr, "dy/dx=") {
		return nil, fmt.Errorf("invalid equation format. Use 'dy/dx = f(x, y)'")
	}
	exprStr := strings.TrimPrefix(equationStr, "dy/dx=")

	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) { return math.Sin(args[0].(float64)), nil },
		"cos": func(args ...interface{}) (interface{}, error) { return math.Cos(args[0].(float64)), nil },
		"tan": func(args ...interface{}) (interface{}, error) { return math.Tan(args[0].(float64)), nil },
		"exp": func(args ...interface{}) (interface{}, error) { return math.Exp(args[0].(float64)), nil },
		"log": func(args ...interface{}) (interface{}, error) { return math.Log(args[0].(float64)), nil },
	}

	parsedExpr, err := govaluate.NewEvaluableExpressionWithFunctions(exprStr, functions)
	if err != nil {
		return nil, fmt.Errorf("error parsing equation: %v", err)
	}
	evaluator := &Evaluator{expr: parsedExpr}

	return func(x, y float64) float64 {
		return evaluator.Evaluate(map[string]interface{}{"x": x, "y": y})
	}, nil
}

// Функция для упрощения и преобразования второго порядка уравнения в вычисляемую функцию
func simplifySecondOrderEquation(equationStr string) (SecondOrderEquation, error) {
	equationStr = strings.ToLower(strings.ReplaceAll(equationStr, " ", ""))
	if !strings.HasPrefix(equationStr, "d2y/dx2=") {
		return nil, fmt.Errorf("invalid equation format. Use 'd2y/dx2 = f(x, y, dy/dx)'")
	}
	exprStr := strings.TrimPrefix(equationStr, "d2y/dx2=")

	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) { return math.Sin(args[0].(float64)), nil },
		"cos": func(args ...interface{}) (interface{}, error) { return math.Cos(args[0].(float64)), nil },
		"tan": func(args ...interface{}) (interface{}, error) { return math.Tan(args[0].(float64)), nil },
		"exp": func(args ...interface{}) (interface{}, error) { return math.Exp(args[0].(float64)), nil },
		"log": func(args ...interface{}) (interface{}, error) { return math.Log(args[0].(float64)), nil },
	}

	parsedExpr, err := govaluate.NewEvaluableExpressionWithFunctions(exprStr, functions)
	if err != nil {
		return nil, fmt.Errorf("error parsing equation: %v", err)
	}
	evaluator := &Evaluator{expr: parsedExpr}

	return func(x, y, dydx float64) float64 {
		return evaluator.Evaluate(map[string]interface{}{"x": x, "y": y, "dydx": dydx})
	}, nil
}

type Result struct {
	X []float64
	Y []float64
}

// Euler method for solving ODE
func eulerMethod(f Equation, x0, y0, h float64, steps int) Result {
	result := Result{X: make([]float64, steps), Y: make([]float64, steps)}
	for i := 0; i < steps; i++ {
		result.X[i] = x0
		result.Y[i] = y0
		y0 += h * f(x0, y0)
		x0 += h
	}
	return result
}

// Runge-Kutta 4th Order Method for solving ODE
func rungeKuttaMethod(f Equation, x0, y0, h float64, steps int) Result {
	result := Result{X: make([]float64, steps), Y: make([]float64, steps)}
	for i := 0; i < steps; i++ {
		result.X[i] = x0
		result.Y[i] = y0
		k1 := h * f(x0, y0)
		k2 := h * f(x0+h/2, y0+k1/2)
		k3 := h * f(x0+h/2, y0+k2/2)
		k4 := h * f(x0+h, y0+k3)
		y0 += (k1 + 2*k2 + 2*k3 + k4) / 6
		x0 += h
	}
	return result
}

// Euler method for solving second-order ODE
func eulerMethodSecondOrder(f SecondOrderEquation, x0, y0, dydx0, h float64, steps int) Result {
	result := Result{X: make([]float64, steps), Y: make([]float64, steps)}
	for i := 0; i < steps; i++ {
		result.X[i] = x0
		result.Y[i] = y0
		y0 += h * dydx0
		dydx0 += h * f(x0, y0, dydx0)
		x0 += h
	}
	return result
}

// Runge-Kutta 4th Order Method for solving second-order ODE
func rungeKuttaMethodSecondOrder(f SecondOrderEquation, x0, y0, dydx0, h float64, steps int) Result {
	result := Result{X: make([]float64, steps), Y: make([]float64, steps)}
	for i := 0; i < steps; i++ {
		result.X[i] = x0
		result.Y[i] = y0
		k1 := h * dydx0
		l1 := h * f(x0, y0, dydx0)
		k2 := h * (dydx0 + l1/2)
		l2 := h * f(x0+h/2, y0+k1/2, dydx0+l1/2)
		k3 := h * (dydx0 + l2/2)
		l3 := h * f(x0+h/2, y0+k2/2, dydx0+l2/2)
		k4 := h * (dydx0 + l3)
		l4 := h * f(x0+h, y0+k3, dydx0+l3)
		y0 += (k1 + 2*k2 + 2*k3 + k4) / 6
		dydx0 += (l1 + 2*l2 + 2*l3 + l4) / 6
		x0 += h
	}
	return result
}

func createChart(result Result) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "ODE Solution"}))

	items := make([]opts.LineData, len(result.X))
	for i := range result.X {
		items[i] = opts.LineData{Value: result.Y[i]}
	}

	line.SetXAxis(result.X).
		AddSeries("y", items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))

	return line
}

func httpServer(result Result) {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		line := createChart(result)
		line.Render(w)
	})

	fmt.Println("Serving chart on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	var method, order string
	fmt.Println("Choose order of ODE (first/second):")
	fmt.Scanln(&order)

	fmt.Println("Choose method (euler/rk4):")
	fmt.Scanln(&method)

	fmt.Println("Enter equation (e.g., dy/dx = sin(x) - y for first order or d2y/dx2 = -sin(x) - y for second order):")
	reader := bufio.NewReader(os.Stdin)
	equationStr, _ := reader.ReadString('\n')
	equationStr = strings.TrimSpace(equationStr)

	x0, y0 := 0.0, 1.0 // Начальные условия
	h := 0.01          // Шаг интегрирования
	steps := 1000      // Количество шагов

	if order == "first" {
		f, err := simplifyEquation(equationStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var result Result
		if method == "euler" {
			result = eulerMethod(f, x0, y0, h, steps)
		} else if method == "rk4" {
			result = rungeKuttaMethod(f, x0, y0, h, steps)
		} else {
			fmt.Println("Invalid method. Use 'euler' or 'rk4'.")
			return
		}
		httpServer(result)
	} else if order == "second" {
		f, err := simplifySecondOrderEquation(equationStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		dydx0 := 0.0 // Начальное значение производной
		var result Result
		if method == "euler" {
			result = eulerMethodSecondOrder(f, x0, y0, dydx0, h, steps)
		} else if method == "rk4" {
			result = rungeKuttaMethodSecondOrder(f, x0, y0, dydx0, h, steps)
		} else {
			fmt.Println("Invalid method. Use 'euler' or 'rk4'.")
			return
		}
		httpServer(result)
	} else {
		fmt.Println("Invalid order. Use 'first' or 'second'.")
		return
	}
}

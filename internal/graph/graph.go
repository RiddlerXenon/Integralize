package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/RiddlerXenon/Integralize/internal/integral/methods"
	"github.com/RiddlerXenon/Integralize/internal/parser"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func createChart(nValues []float64, errorsLeft, errorsRight, errorsMid []float64) *charts.Line {
	graph := charts.NewLine()

	xValues := make([]string, len(nValues))
	for i, n := range nValues {
		xValues[i] = fmt.Sprintf("%d", int(n))
	}

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "\n\nПогрешности",
			Subtitle: "Сравнение известных методов интегрирования",
		}),
		charts.WithGridOpts(opts.Grid{
			Top: "20%",
		}),
	)

	graph.SetXAxis(xValues).
		AddSeries("Left Rectangle", generateLineItems(errorsLeft)).
		AddSeries("Right Rectangle", generateLineItems(errorsRight)).
		AddSeries("Midpoint Rectangle", generateLineItems(errorsMid))
		// AddSeries("Trapezional", generateLineItems(errorsTrapez)).
		// AddSeries("Simpson", generateLineItems(errorsSimps)).
		// AddSeries("Monte-Carlo", generateLineItems(errorsMonte)).
		// AddSeries("Gauss", generateLineItems(errorsGauss)).
		// AddSeries("Chebyshev", generateLineItems(errorsCheb))

	return graph
}

// Генерация точек графика
func generateLineItems(data []float64) []opts.LineData {
	items := make([]opts.LineData, len(data))
	for i, v := range data {
		items[i] = opts.LineData{Value: v}
	}
	return items
}

func main() {
	a, b := 0.0, math.Pi // Пример интегрирования функции sin(x) от 0 до Pi
	fun, _ := parser.ParseStr("(sin(X) + exp(X))/(2^3)")
	fmt.Println(fun(10))
	nValues := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	trueValue := 2.0

	errorsLeft := make([]float64, len(nValues))
	errorsRight := make([]float64, len(nValues))
	errorsMid := make([]float64, len(nValues))
	// errorsTrapez := make([]float64, len(nValues))
	// errorsSimps := make([]float64, len(nValues))
	// errorsMonte := make([]float64, len(nValues))
	// errorsGauss := make([]float64, len(nValues))
	// errorsCheb := make([]float64, len(nValues))

	for i, n := range nValues {
		leftResult, err := methods.LeftRectangleMethod(a, b, n, fun)
		if err != nil {
			log.Fatalf("Error in LeftRectangleMethod: %v", err)
		}
		errorsLeft[i] = math.Abs(trueValue - leftResult)

		rightResult, err := methods.RightRectangleMethod(a, b, n, fun)
		if err != nil {
			log.Fatalf("Error in RightRectangleMethod: %v", err)
		}
		errorsRight[i] = math.Abs(trueValue - rightResult)

		midResult, err := methods.MidpointRectangleMethod(a, b, n, fun)
		if err != nil {
			log.Fatalf("Error in MidpointRectangleMethod: %v", err)
		}
		errorsMid[i] = math.Abs(trueValue - midResult)

		// 	TrapezResult, err := methods.TrapezoidMethod(a, b, n, expr)
		// 	if err != nil {
		// 		log.Fatalf("Error in TrapezMethod: %v", err)
		// 	}
		// 	errorsTrapez[i] = math.Abs(trueValue - TrapezResult)

		// 	SimpsonResult, err := methods.SimpsonMethod(a, b, n, expr)
		// 	if err != nil {
		// 		log.Fatalf("Error in SimpsonMethod: %v", err)
		// 	}
		// 	errorsSimps[i] = math.Abs(trueValue - SimpsonResult)

		// 	MonteResult, err := methods.MonteCarloMethod(a, b, n, expr)
		// 	if err != nil {
		// 		log.Fatalf("Error in MonteCarloMethod: %v", err)
		// 	}
		// 	errorsMonte[i] = math.Abs(trueValue - MonteResult)

		// 	GaussResult, err := methods.GaussQuadrature(a, b, int(n), expr)
		// 	if err != nil {
		// 		log.Fatalf("Error in GaussQuadrature: %v", err)
		// 	}
		// 	errorsGauss[i] = math.Abs(trueValue - GaussResult)

		// 	ChebyshevResult, err := methods.ChebyshevQuadrature(a, b, int(n), expr)
		// 	if err != nil {
		// 		log.Fatalf("Error in ChebyshevQuadrature: %v", err)
		// 	}
		// 	errorsCheb[i] = math.Abs(trueValue - ChebyshevResult)
	}

	graph := createChart(nValues, errorsLeft, errorsRight, errorsMid)

	f, _ := os.Create("accuracy_chart.html")
	graph.Render(f)

	fmt.Println("График сохранен в файл accuracy_chart.html")
}

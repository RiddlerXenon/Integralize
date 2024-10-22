package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/RiddlerXenon/Integralize/internal/integral/methods"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func createChart(nValues []float64, errorsLeft, errorsRight, errorsMid []float64) *charts.Line {
	graph := charts.NewLine()

	xValues := make([]string, len(nValues))
	for i, n := range nValues {
		xValues[i] = fmt.Sprintf("%f", n)
	}

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Погрешности",
			Subtitle: "Сравнение левого, правого и среднего метода",
		}),
	)

	graph.SetXAxis(xValues).
		AddSeries("Left Rectangle", generateLineItems(errorsLeft)).
		AddSeries("Right Rectangle", generateLineItems(errorsRight)).
		AddSeries("Midpoint Rectangle", generateLineItems(errorsMid))

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
	expr := "sin(x)"
	nValues := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	trueValue := 2.0

	errorsLeft := make([]float64, len(nValues))
	errorsRight := make([]float64, len(nValues))
	errorsMid := make([]float64, len(nValues))

	for i, n := range nValues {
		leftResult, err := methods.LeftRectangleMethod(a, b, n, expr)
		if err != nil {
			log.Fatalf("Error in LeftRectangleMethod: %v", err)
		}
		errorsLeft[i] = math.Abs(trueValue - leftResult)

		rightResult, err := methods.RightRectangleMethod(a, b, n, expr)
		if err != nil {
			log.Fatalf("Error in RightRectangleMethod: %v", err)
		}
		errorsRight[i] = math.Abs(trueValue - rightResult)

		midResult, err := methods.MidpointRectangleMethod(a, b, n, expr)
		if err != nil {
			log.Fatalf("Error in MidpointRectangleMethod: %v", err)
		}
		errorsMid[i] = math.Abs(trueValue - midResult)
	}

	graph := createChart(nValues, errorsLeft, errorsRight, errorsMid)

	f, _ := os.Create("accuracy_chart.html")
	graph.Render(f)

	fmt.Println("График сохранен в файл accuracy_chart.html")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/RiddlerXenon/Integralize/internal/differential"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	params := differential.Parameters{
		Alpha: 0.1,
		Beta:  0.02,
		Gamma: 0.3,
		Delta: 0.01,
		Step:  0.1,
		Steps: 1000,
		Prey:  40,
		Pred:  9,
	}
	preyEuler, predEuler, errEuler := differential.EulerMethod(params)

	if errEuler != nil {
		fmt.Println("Error in Euler method: ", errEuler)
		return
	}

	preyRK, predRK, errRK := differential.RungeKuttaMethod(params)
	if errRK != nil {
		fmt.Println("Error in Runge-Kutta method: ", errRK)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		eulerChart := createChart("Lotka-Volterra Model (Euler Method)", preyEuler, predEuler, params.Steps)
		rkChart := createChart("Lotka-Volterra Model (Runge-Kutta Method)", preyRK, predRK, params.Steps)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<html><body>")
		eulerChart.Render(w)
		fmt.Fprint(w, "<br>")
		rkChart.Render(w)
		fmt.Fprint(w, "</body></html>")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func createChart(title string, prey, pred []float64, steps int) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: title,
	}))

	xAxis := make([]int, steps)
	for i := 0; i < steps; i++ {
		xAxis[i] = i
	}

	line.SetXAxis(xAxis).
		AddSeries("Prey", generateLineItems(prey)).
		AddSeries("Predator", generateLineItems(pred))

	return line
}

func generateLineItems(data []float64) []opts.LineData {
	items := make([]opts.LineData, len(data))
	for i, d := range data {
		items[i] = opts.LineData{Value: d}
	}
	return items
}

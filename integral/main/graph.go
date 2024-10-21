package main

import (
	"fmt"
	"math"

	"diplo/integral/methods"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Функция для вычисления истинного значения интеграла
func trueValue() float64 {
	return 2.0 // Точное значение интеграла sin(x) от 0 до π
}

// Функция для построения графика
func plotMethods(method string, start, end int) {
	// Создаем новый график
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Точность метода %s", method)
	p.X.Label.Text = "Количество узлов"
	p.Y.Label.Text = "Ошибка"

	// Определяем точки для графика
	points := make(plotter.XYs, end-start+1)

	for n := start; n <= end; n++ {
		var result float64
		switch method {
		case "rectangle":
			result = methods.LeftRectangleMethod(0, math.Pi, n) // Можно изменить на любой метод
		case "trapezoid":
			result = methods.TrapezoidMethod(0, math.Pi, n)
		case "simpson":
			result = methods.SimpsonMethod(0, math.Pi, n)
		case "montecarlo":
			result = methods.MonteCarloMethod(0, math.Pi, n)
		default:
			fmt.Println("Неверный метод")
			return
		}
		error := math.Abs(result - trueValue())
		points[n-start].X = float64(n)
		points[n-start].Y = error
	}

	// Создаем линию для графика
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Сохраняем график в файл
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s_accuracy.png", method)); err != nil {
		panic(err)
	}
}

func main() {
	var method string
	var start, end int

	fmt.Println("Введите метод (rectangle, trapezoid, simpson, montecarlo, newtoncotes):")
	fmt.Scan(&method)
	fmt.Println("Введите диапазон узлов (начало и конец):")
	fmt.Scan(&start, &end)

	plotMethods(method, start, end)
	fmt.Println("График успешно создан!")
}

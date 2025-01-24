package main

import (
	"fmt"
	"log"

	"github.com/RiddlerXenon/Integralize/internal/parser"
)

func eulerMethod(f func(float64, float64) float64, y0, t0, tMax, h float64) ([]float64, []float64) {
	//Разбиваем на шаги
	nSteps := int((tMax - t0) / h)

	/* Так как изначально диффур принимает вид как dy/dt,
	то логично использовать X = t, Y = y */
	t := make([]float64, nSteps+1)
	y := make([]float64, nSteps+1)

	t[0] = t0 //Начальное время
	y[0] = y0 //y(0) = Введенному значению

	/* Воспользуемся формулой
	y[i + 1] = y[i] - h * k * y[i]
	t[i + 1] = t[i] + h
	где k - коэффициент, y0 - начальное значение
	tMax - максимальное время, h - шаг */
	for i := 0; i < nSteps; i++ {
		t[i+1] = t[i] + h
		y[i+1] = y[i] + h*f(t[i], y[i])
	}

	return t, y
}

func rungeKutte(f func(float64, float64) float64, y0, t0, tMax, h float64) ([]float64, []float64) {
	nSteps := int((tMax - t0) / h)
	t := make([]float64, nSteps+1)
	y := make([]float64, nSteps+1)

	t[0] = t0
	y[0] = y0

	for i := 0; i < nSteps; i++ {
		t[i+1] = t[i] + h

		k1 := h * f(t[i], y[i])
		k2 := h * f(t[i]+h/2, y[i]+k1/2)
		k3 := h * f(t[i]+h/2, y[i]+k2/2)
		k4 := h * f(t[i]+h, y[i]+k3)

		y[i+1] = y[i] + (k1+2*k2+2*k3+k4)/6
	}

	return t, y
}

func main() {
	eqStr := "X*X + sqrt(Y)"

	parsFunc, err := parser.ParseStrDiffEq(eqStr)
	if err != nil {
		log.Fatalf("Ошибка парсинга: %v", err)
	}

	//Начальные условия для примера
	y0 := 1.0
	t0 := 0.0
	tMax := 2.0
	h := 0.1

	t, y := eulerMethod(parsFunc, y0, t0, tMax, h)
	t1, y1 := rungeKutte(parsFunc, y0, t0, tMax, h)

	fmt.Printf("Метод Эйлера\n")
	for i := range t {
		fmt.Printf("t=%.2f, y=%.4f\n", t[i], y[i])
	}

	fmt.Printf("\nМетод Рунге-Кутты\n")
	for i := range t {
		fmt.Printf("t=%.2f, y=%.4f\n", t1[i], y1[i])
	}
}

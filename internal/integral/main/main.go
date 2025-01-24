package main

import (
	"fmt"
	"log"
	"math"

	"github.com/RiddlerXenon/Integralize/internal/integral/methods"
	"github.com/RiddlerXenon/Integralize/internal/parser"
)

func main() {

	expr, err := parser.ParseStrInt("sin(X)") // Можно заменить на любое выражение с "X"
	if err != nil {
		log.Fatalf("Ошибка парсинга выражения: %v", err)
	}

	a := 0.0
	b := math.Pi
	n := 10.0

	leftRes, err := methods.LeftRectangleMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in LeftRectangleMethod: %v", err)
	}
	fmt.Printf("Решение через левых прямоугольников: %f\n", leftRes)

	RightRes, err := methods.RightRectangleMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in RightRectangleMethod: %v", err)
	}
	fmt.Printf("Решение через правых прямоугольников: %f\n", RightRes)

	MidRes, err := methods.MidpointRectangleMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in MidpointRectangleMethod: %v", err)
	}
	fmt.Printf("Решение через центральных прямоугольников: %f\n", MidRes)

	TrapezResult, err := methods.TrapezoidMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in TrapezMethod: %v", err)
	}
	fmt.Printf("Решение через метод трапеций: %f\n", TrapezResult)

	SimpsonResult, err := methods.SimpsonMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in SimpsonMethod: %v", err)
	}
	fmt.Printf("Решение через метод Симпсона: %f\n", SimpsonResult)

	MonteResult, err := methods.MonteCarloMethod(a, b, n, expr)
	if err != nil {
		log.Fatalf("Error in MonteCarloMethod: %v", err)
	}
	fmt.Printf("Решение через метод Монте Карло: %f\n", MonteResult)

	GaussResult, err := methods.GaussQuadrature(a, b, int(n), expr)
	if err != nil {
		log.Fatalf("Error in GaussQuadrature: %v", err)
	}
	fmt.Printf("Решение через метод Гаусса: %f\n", GaussResult)

	ChebyshevResult, err := methods.ChebyshevQuadrature(a, b, int(n), expr)
	if err != nil {
		log.Fatalf("Error in ChebyshevQuadrature: %v", err)
	}
	fmt.Printf("Решение через метод Чебышева: %f\n", ChebyshevResult)

}

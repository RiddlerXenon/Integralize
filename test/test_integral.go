package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RiddlerXenon/Integralize/internal/integral"
	"github.com/RiddlerXenon/Integralize/internal/parser"
)

// Определение типа для методов интегрирования
type IntegralMethod struct {
	Name        string
	Description string
	Method      func(float64, float64, float64, func(map[string]float64) float64) (float64, error)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Тестирование методов численного интегрирования ===")
	fmt.Println("Программа рассчитает интеграл всеми доступными методами одновременно")

	// Получаем выражение для интегрирования от пользователя
	fmt.Print("Введите подынтегральное выражение в формате LaTeX (например, x^2 + 2*x + 1): ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	// Парсим выражение
	expressionFunc, err := parser.PrepareLatexExpression(expression)
	if err != nil {
		fmt.Printf("Ошибка при парсинге выражения: %v\n", err)
		return
	}

	// Получаем границы интегрирования
	fmt.Print("Введите нижнюю границу интегрирования (a): ")
	aStr, _ := reader.ReadString('\n')
	a, err := strconv.ParseFloat(strings.TrimSpace(aStr), 64)
	if err != nil {
		fmt.Printf("Ошибка при чтении нижней границы: %v\n", err)
		return
	}

	fmt.Print("Введите верхнюю границу интегрирования (b): ")
	bStr, _ := reader.ReadString('\n')
	b, err := strconv.ParseFloat(strings.TrimSpace(bStr), 64)
	if err != nil {
		fmt.Printf("Ошибка при чтении верхней границы: %v\n", err)
		return
	}

	fmt.Print("Введите количество точек разбиения (n): ")
	nStr, _ := reader.ReadString('\n')
	n, err := strconv.ParseFloat(strings.TrimSpace(nStr), 64)
	if err != nil {
		fmt.Printf("Ошибка при чтении количества точек: %v\n", err)
		return
	}

	// Создаем список доступных методов интегрирования
	methods := []IntegralMethod{
		{"1", "Метод левых прямоугольников", integral.LeftRectangle},
		{"2", "Метод правых прямоугольников", integral.RightRectangle},
		{"3", "Метод средних прямоугольников", integral.MidpointRectangle},
		{"4", "Метод трапеций", integral.Trapezoidal},
		{"5", "Метод Симпсона", integral.Simpson},
		{"6", "Метод Монте-Карло", integral.MonteCarlo},
		{"7", "Метод Чебышева", integral.Chebyshev},
		{"8", "Метод Гаусса", integral.GaussLegendre},
	}

	// Выводим заголовок результатов
	fmt.Printf("\nРезультаты вычисления интеграла от %f до %f для выражения \"%s\":\n", a, b, expression)
	fmt.Printf("Количество разбиений: %d\n\n", int(n))

	fmt.Println("+---------------------------------+----------------+----------------+")
	fmt.Println("| Метод                          | Результат      | Время (мс)     |")
	fmt.Println("+---------------------------------+----------------+----------------+")

	// Вычисляем и выводим результаты для всех методов
	for _, method := range methods {
		// Замеряем время выполнения
		startTime := time.Now()
		result, err := method.Method(a, b, n, expressionFunc)
		duration := time.Since(startTime).Milliseconds()

		// Форматируем результат или сообщение об ошибке
		resultStr := "ошибка"
		if err == nil {
			resultStr = fmt.Sprintf("%.8f", result)
		}

		// Выводим результат в табличном виде
		fmt.Printf("| %-31s | %-14s | %-14d |\n", method.Description, resultStr, duration)
	}
	fmt.Println("+---------------------------------+----------------+----------------+")

	// Проверка с другими значениями n для анализа сходимости по выбранному пользователем методу
	fmt.Println("\nХотите проверить сходимость определённого метода с разными значениями n? (да/нет)")
	choiceStr, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(strings.ToLower(choiceStr))

	if choice == "да" || choice == "y" || choice == "yes" {
		// Выводим список методов для выбора
		fmt.Println("\nВыберите метод для проверки сходимости:")
		for _, method := range methods {
			fmt.Printf("%s. %s\n", method.Name, method.Description)
		}

		// Запрашиваем выбор метода
		fmt.Print("\nВведите номер метода: ")
		methodChoiceStr, _ := reader.ReadString('\n')
		methodChoice := strings.TrimSpace(methodChoiceStr)

		var selectedMethod func(float64, float64, float64, func(map[string]float64) float64) (float64, error)
		var methodName string

		// Находим выбранный метод
		for _, method := range methods {
			if method.Name == methodChoice {
				selectedMethod = method.Method
				methodName = method.Description
				break
			}
		}

		if selectedMethod != nil {
			fmt.Printf("\nПроверка сходимости для метода \"%s\":\n\n", methodName)
			fmt.Println("+--------+----------------+----------------+")
			fmt.Println("| n      | Результат      | Время (мс)     |")
			fmt.Println("+--------+----------------+----------------+")

			// Тестируем с разными значениями n
			testValues := []float64{n, n * 2, n * 5, n * 10, n * 100}
			for _, tn := range testValues {
				startTime := time.Now()
				result, err := selectedMethod(a, b, tn, expressionFunc)
				duration := time.Since(startTime).Milliseconds()

				resultStr := "ошибка"
				if err == nil {
					resultStr = fmt.Sprintf("%.8f", result)
				}

				fmt.Printf("| %-6d | %-14s | %-14d |\n", int(tn), resultStr, duration)
			}
			fmt.Println("+--------+----------------+----------------+")
		} else {
			fmt.Println("Выбран неверный метод")
		}
	}

	// Предлагаем вычислить аналитический интеграл, если известен
	fmt.Println("\nХотите ввести известное аналитическое решение для сравнения? (да/нет)")
	analyticChoiceStr, _ := reader.ReadString('\n')
	analyticChoice := strings.TrimSpace(strings.ToLower(analyticChoiceStr))

	if analyticChoice == "да" || analyticChoice == "y" || analyticChoice == "yes" {
		fmt.Print("Введите значение аналитического решения: ")
		analyticStr, _ := reader.ReadString('\n')
		analytic, err := strconv.ParseFloat(strings.TrimSpace(analyticStr), 64)

		if err == nil {
			fmt.Println("\nСравнение с аналитическим решением:")
			fmt.Println("+--------------------------------+--------------+--------------+")
			fmt.Println("| Метод                         | Результат    | Погрешность  |")
			fmt.Println("+--------------------------------+--------------+--------------+")

			for _, method := range methods {
				result, err := method.Method(a, b, n, expressionFunc)

				resultStr := "ошибка"
				errorStr := "----"
				if err == nil {
					resultStr = fmt.Sprintf("%.8f", result)
					errorStr = fmt.Sprintf("%.8f", math.Abs(result-analytic))
				}

				fmt.Printf("| %-30s | %-12s | %-12s |\n", method.Description, resultStr, errorStr)
			}
			fmt.Printf("| %-30s | %-12.8f | %-12s |\n", "Аналитическое решение", analytic, "----")
			fmt.Println("+--------------------------------+--------------+--------------+")
		} else {
			fmt.Println("Ошибка ввода аналитического решения")
		}
	}

	fmt.Println("\nТестирование завершено.")
}
